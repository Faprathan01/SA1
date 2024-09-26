import React, { useState } from 'react';
import DatePicker from 'react-datepicker';
import 'react-datepicker/dist/react-datepicker.css';

interface PaymentMethodProps {
  selectedMethod: string;
  onMethodChange: (method: string) => void;
}

const PaymentMethod: React.FC<PaymentMethodProps> = ({ selectedMethod, onMethodChange }) => {
  const [expiryDate, setExpiryDate] = useState<Date | null>(null);
  const [paypalEmail, setPaypalEmail] = useState<string>('');
  const [paypalPassword, setPaypalPassword] = useState<string>('');
  const [promptPayPhone, setPromptPayPhone] = useState<string>('');

  // การสร้าง URL สำหรับ QR Code
  const promptPayQRCodeURL = "https://png.pngtree.com/png-clipart/20220729/original/pngtree-qr-code-png-image_8438558.png"; // เปลี่ยน URL นี้ให้เป็น URL ของ QR Code จริง

  return (
    <div className="bg-gray-100 p-8 rounded-xl shadow-lg flex-1 mb-6 md:mb-0 md:mr-6">
      <h2 className="text-4xl font-semibold mb-8 text-gray-900">Select Payment Method</h2>

      <div className="flex justify-start mb-6 space-x-4">
        <button
          className={`py-4 px-6 rounded-md shadow-md focus:outline-none ${selectedMethod === 'Credit Card' ? 'bg-lime-400 text-white' : 'bg-gray-200'}`}
          onClick={() => onMethodChange('Credit Card')}
        >
          Credit Card
        </button>
        <button
          className={`py-4 px-6 rounded-md shadow-md focus:outline-none ${selectedMethod === 'PromptPay' ? 'bg-lime-400 text-white' : 'bg-gray-200'}`}
          onClick={() => onMethodChange('PromptPay')}
        >
          PromptPay
        </button>
        <button
          className={`py-4 px-6 rounded-md shadow-md focus:outline-none ${selectedMethod === 'PayPal' ? 'bg-lime-400 text-white' : 'bg-gray-200'}`}
          onClick={() => onMethodChange('PayPal')}
        >
          PayPal
        </button>
      </div>

      {selectedMethod === 'Credit Card' && (
        <>
          <div className="mb-6">
            <label className="block text-gray-700 font-semibold mb-3 text-xl" htmlFor="cardHolder">
              Name on Card
            </label>
            <input type="text" id="cardHolder" className="w-full p-4 rounded-2xl border border-gray-300 text-lg text-black" placeholder="Enter card holder name" />
          </div>
          <div className="mb-6">
            <label className="block text-gray-700 font-semibold mb-3 text-xl" htmlFor="cardNumber">
              Card Number
            </label>
            <input type="text" id="cardNumber" className="w-full p-4 rounded-2xl border border-gray-300 text-lg text-black" placeholder="0000 - 0000 - 0000 - 0000" />
          </div>
          <div className="flex justify-between mb-4">
            <div className="w-1/2 mr-4">
              <label className="block text-gray-700 font-semibold mb-3 text-xl" htmlFor="expiryDate">
                Expiry Date
              </label>
              <DatePicker selected={expiryDate} onChange={(date: Date | null) => setExpiryDate(date)} dateFormat="dd/MM/yyyy" placeholderText="Select Date" className="w-full p-4 rounded-2xl border border-gray-300 text-lg text-black" />
            </div>
            <div className="w-1/2 ml-4">
              <label className="block text-gray-700 font-semibold mb-3 text-xl" htmlFor="cvvCode">
                CVV
              </label>
              <input type="text" id="cvvCode" className="w-full p-4 rounded-2xl border border-gray-300 text-lg text-black" placeholder="CVV" />
            </div>
          </div>
        </>
      )}

      {selectedMethod === 'PromptPay' && (  
        <>
          <div className="mb-6">
            <label className="block text-gray-700 font-semibold mb-3 text-xl" htmlFor="promptPayPhone">
              Phone Number for PromptPay
            </label>
            <input
              type="tel"
              id="promptPayPhone"
              className="w-full p-4 rounded-2xl border border-gray-300 text-lg text-black"
              placeholder="Enter phone number"
              value={promptPayPhone}
              onChange={(e) => setPromptPayPhone(e.target.value)}
            />
          </div>
          <div className="mb-6">
            <h3 className="block text-gray-700 font-semibold mb-3 text-xl">
              PromptPay QR Code
            </h3>
            {/* แสดง QR Code */}
            <div className="w-full p-4 rounded-2xl border border-gray-300 flex justify-center items-center">
              <img src={promptPayQRCodeURL} alt="PromptPay QR Code" className="w-48 h-48" />
            </div>
          </div>
        </>
      )}

      {selectedMethod === 'PayPal' && (
        <>
          <div className="mb-6">
            <label className="block text-gray-700 font-semibold mb-3 text-xl" htmlFor="paypalEmail">
              PayPal Email
            </label>
            <input
              type="email"
              id="paypalEmail"
              className="w-full p-4 rounded-2xl border border-gray-300 text-lg text-black"
              placeholder="Enter PayPal email"
              value={paypalEmail}
              onChange={(e) => setPaypalEmail(e.target.value)}
            />
          </div>
          <div className="mb-6">
            <label className="block text-gray-700 font-semibold mb-3 text-xl" htmlFor="paypalPassword">
              PayPal Password
            </label>
            <input
              type="password"
              id="paypalPassword"
              className="w-full p-4 rounded-2xl border border-gray-300 text-lg text-black"
              placeholder="Enter PayPal password"
              value={paypalPassword}
              onChange={(e) => setPaypalPassword(e.target.value)}
            />
          </div>
        </>
      )}
    </div>
  );
};

export default PaymentMethod;
