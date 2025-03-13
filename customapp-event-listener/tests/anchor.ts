import * as anchor from "@coral-xyz/anchor";
import * as anchor from "@coral-xyz/anchor";
import { Program } from "@coral-xyz/anchor";
import { Event } from "../target/types/event";
import type { Event } from "../target/types/event";
 
describe("event", () => {
  // Configure the client to use the local cluster
  anchor.setProvider(anchor.AnchorProvider.env());

  const program = anchor.workspace.Event as anchor.Program<Event>;
  
  // Configure the client to use the local cluster.
  anchor.setProvider(anchor.AnchorProvider.env());
 
  const program = anchor.workspace.Event as Program<Event>;
 
  it("Emits custom event", async () => {
    // Set up listener before sending transaction
    const listenerId = program.addEventListener("CustomEvent", event => {
      // Do something with the event data
      console.log("Event Data:", event);
    });
 
    // Message to be emitted in the event
    const message = "Hello, Solana!";
    // Send transaction
    await program.methods.emitEvent(message).rpc();
 
    // Remove listener
    await program.removeEventListener(listenerId);
  });
});