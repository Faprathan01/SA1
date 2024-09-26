export interface PaymentInterface {
    PaymentID?: number;
    MethodName?: string; // เช่น 'Credit Card' หรือ 'PromptPay'
    Status?: string;
    Amount?: number;
}   
