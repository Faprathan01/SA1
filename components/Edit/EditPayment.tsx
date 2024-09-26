import React, { useState, useEffect } from "react";


// ตัวอย่างฟังก์ชันการดึงข้อมูลการชำระเงินที่มีอยู่
const PaymentData = (paymentId: string) => {
    return {
        paymentId: paymentId,
        methodName: "Paypal",
        amount: "1999.00",
        memberId: "1", // สมมุติเป็นสมาชิกที่มี id 1
        packageId: "3", // สมมุติเป็นแพ็คเกจที่มี id 3
    };
};

interface PaymentData {
    paymentId: string;
    methodName: string;
    amount: string;
    memberId: string;
    packageId: string;
}

const EditPayment: React.FC<{ paymentId: string }> = ({ paymentId }) => {
    const [formData, setFormData] = useState<PaymentData>({
        paymentId: "",
        methodName: "",
        amount: "",
        memberId: "",
        packageId: "",
    });

    // ดึงข้อมูลการชำระเงินปัจจุบัน
    useEffect(() => {
        const data = PaymentData(paymentId);
        setFormData(data);
    }, [paymentId]);

    const handleInputChange = (
        e: React.ChangeEvent<HTMLInputElement | HTMLSelectElement | HTMLTextAreaElement>,
        field: keyof typeof formData
    ) => {
        setFormData({
            ...formData,
            [field]: e.target.value,
        });
    };

    const handleSubmit = (e: React.FormEvent) => {
        e.preventDefault();
        console.log("Updated Payment Data:", formData);
        // คุณสามารถส่งข้อมูลที่แก้ไขแล้วไปยัง API ที่นี่
    };

    return (
        <div className="w-screen max-w-7xl mx-auto p-12 bg-gray-800 rounded-lg shadow-lg mt-10">
            
            <form onSubmit={handleSubmit}>
                <div className="grid grid-cols-1 md:grid-cols-2 gap-10">
                    {/* Payment Method Name */}
                    <div>
                        <label className="block text-green-400 mb-2 text-3xl">Payment Method</label>
                        <input
                            id="methodName"
                            name="methodName"
                            type="text"
                            required
                            placeholder="Enter payment method"
                            value={formData.methodName}
                            onChange={(e) => handleInputChange(e, "methodName")}
                            className="w-full p-3 bg-white text-black text-2xl rounded-lg focus:outline-none focus:ring-2 focus:ring-green-300"
                        />
                    </div>

                    {/* Payment Amount */}
                    <div>
                        <label className="block text-green-400 mb-2 text-3xl">Amount</label>
                        <input
                            id="amount"
                            name="amount"
                            type="number"
                            required
                            placeholder="Enter amount"
                            value={formData.amount}
                            onChange={(e) => handleInputChange(e, "amount")}
                            className="w-full p-3 bg-white text-black text-2xl rounded-lg focus:outline-none focus:ring-2 focus:ring-green-300"
                        />
                    </div>

                    {/* Member ID */}
                    <div>
                        <label className="block text-green-400 mb-2 text-3xl">Member</label>
                        <select
                            id="memberId"
                            name="memberId"
                            required
                            value={formData.memberId}
                            onChange={(e) => handleInputChange(e, "memberId")}
                            className="w-full p-3 bg-white text-black text-2xl rounded-lg focus:outline-none focus:ring-2 focus:ring-green-300"
                        >
                            {/* ตัวอย่างข้อมูลสมาชิก */}
                            <option value="1">Member 1</option>
                            <option value="2">Member 2</option>
                            <option value="3">Member 3</option>
                        </select>
                    </div>

                    {/* Package ID */}
                    <div>
                        <label className="block text-green-400 mb-2 text-3xl">Package</label>
                        <select
                            id="packageId"
                            name="packageId"
                            required
                            value={formData.packageId}
                            onChange={(e) => handleInputChange(e, "packageId")}
                            className="w-full p-3 bg-white text-black text-2xl rounded-lg focus:outline-none focus:ring-2 focus:ring-green-300"
                        >
                            {/* ตัวอย่างข้อมูลแพ็คเกจ */}
                            <option value="1">Package 1</option>
                            <option value="2">Package 2</option>
                            <option value="3">Package 3</option>
                        </select>
                    </div>
                </div>

                {/* Buttons */}
                <div className="flex justify-end space-x-4 mt-8">
                    <button
                        type="button"
                        className="bg-red hover:bg-rose-500 text-black font-semibold py-4 px-8 rounded-lg focus:outline-none focus:ring-2 focus:ring-red-500 transition-all duration-300 ease-in-out text-2xl"
                    >
                        Cancel
                    </button>
                    <button
                        type="submit"
                        className="bg-green-400 hover:bg-green-500 text-black font-semibold py-4 px-8 rounded-lg focus:outline-none focus:ring-2 focus:ring-green-500 transition-all duration-300 ease-in-out text-2xl"
                    >
                        Update Payment
                    </button>
                </div>
            </form>
        </div>
    );
};

export default EditPayment;
