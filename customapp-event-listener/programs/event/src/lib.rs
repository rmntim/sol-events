use anchor_lang::prelude::*;
 
declare_id!("6QSaTnP3biE27vP1gRFqmHr7hzpp3VRTdFnEonuMtgWw");
 
#[program]
pub mod event {
    use super::*;

    pub fn say_hello(_ctx: Context<EmitEvent>) -> Result<()> {
        msg!("hello solana!");
        Ok(())
    }
 
    pub fn emit_event(_ctx: Context<EmitEvent>, input: String) -> Result<()> {
        emit!(CustomEvent { message: input.clone() });
        emit!(CustomEvent { message: input.clone() });
        emit!(CustomEvent { message: input });
        Ok(())
    }
}
 
#[derive(Accounts)]
pub struct EmitEvent {}

#[event]
pub struct CustomEvent {
    pub message: String,
}