import * as anchor from "@coral-xyz/anchor";
import type { Event } from "../target/types/event";

// Configure the client to use the local cluster
anchor.setProvider(anchor.AnchorProvider.env());

const program = anchor.workspace.Event as anchor.Program<Event>;

const message = "Hello, Solana!";
// Send transaction
console.log(program.programId.toBase58());
await program.methods.sayHello().rpc();
await program.methods.emitEvent(message).rpc();