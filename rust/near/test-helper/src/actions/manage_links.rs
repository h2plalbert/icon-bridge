use crate::types::{Bmc, Context, Contract};
use crate::{invoke_call, invoke_view};
use duplicate::duplicate;
use near_primitives::types::Gas;

#[duplicate(
    contract_type;
    [ Bmc ];
)]

impl Contract<'_, contract_type> {
    pub fn add_link(&self, mut context: Context, gas: Gas) -> Context {
        invoke_call!(self, context, "add_link", method_params, None, Some(gas));
        context
    }

    pub fn remove_link(&self, mut context: Context) -> Context {
        invoke_call!(self, context, "remove_link", method_params);
        context
    }

    pub fn get_links(&self, mut context: Context) -> Context {
        invoke_view!(self, context, "get_links");
        context
    }

    pub fn set_link(&self, mut context: Context, gas: Gas) -> Context {
        invoke_call!(self, context, "set_link", method_params, None, Some(gas));
        context
    }

    pub fn get_status(&self, mut context: Context) -> Context {
        invoke_view!(self, context, "get_status", method_params);
        context
    }
}
