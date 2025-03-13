package main

import (
	"context"
	"flag"
	"log"
	"sol-events/generated/idl"
	"time"

	solana "github.com/gagliardetto/solana-go"
	rpc "github.com/gagliardetto/solana-go/rpc"
	"golang.org/x/time/rate"
)

func main() {
	programIDString := flag.String("program", "6QSaTnP3biE27vP1gRFqmHr7hzpp3VRTdFnEonuMtgWw", "program id")
	pollingInterval := flag.Duration("interval", 10*time.Second, "polling interval")
	endpoint := flag.String("endpoint", rpc.DevNet_RPC, "endpoint")
	flag.Parse()

	// Choose your endpoint (e.g., MainNetBeta, DevNet, or a custom endpoint)
	client := rpc.NewWithCustomRPCClient(rpc.NewWithLimiter(*endpoint, rate.Every(10*time.Second), 20))

	// Set the program ID (smart contract public key) you want to monitor.
	// Replace "YourProgramID" with your actual program ID.
	programID, err := solana.PublicKeyFromBase58(*programIDString)
	if err != nil {
		log.Fatalf("Failed to parse program ID: %v", err)
	}

	// Keep track of the most recent processed signature.
	var lastSignature string

	// Poll every 10 seconds.
	ticker := time.NewTicker(*pollingInterval)
	defer ticker.Stop()

	for range ticker.C {
		// Get the latest signatures for the given program address.
		sigs, err := client.GetSignaturesForAddress(
			context.Background(),
			programID,
		)
		if err != nil {
			log.Printf("Error retrieving signatures: %v", err)
			continue
		}

		// Process each signature;
		// they are returned in most-recent-first order.
		for _, sigInfo := range sigs {
			// Skip signatures that have already been processed.
			if sigInfo.Signature.String() == lastSignature {
				break
			}

			// Retrieve the full transaction details.
			tx, err := client.GetTransaction(
				context.Background(),
				sigInfo.Signature,
				&rpc.GetTransactionOpts{
					Encoding: solana.EncodingBase64,
				},
			)
			if err != nil {
				log.Printf("Error getting transaction for %s: %v",
					sigInfo.Signature, err)
				continue
			}

			// Check if the transaction contains logs.
			if tx != nil && tx.Meta != nil && len(tx.Meta.LogMessages) > 0 {
				events, err := idl.DecodeEvents(
					tx,
					programID,
					func(altAddresses []solana.PublicKey) (tables map[solana.PublicKey]solana.PublicKeySlice, err error) {
						for _, altAddress := range altAddresses {
							log.Printf("Alt address: %s", altAddress.String())
						}
						return nil, nil
					},
				)
				if err != nil {
					log.Printf("Error decoding events: %v", err)
					continue
				}

				for _, event := range events {
					log.Printf("Event: %s", event.Name)
					log.Printf("Data: %#+v", event.Data)
				}
			}
		}

		// Update the last processed signature (if we got any signatures).
		if len(sigs) > 0 {
			lastSignature = sigs[0].Signature.String()
		}
	}
}
