import React from "react";
import Nav from "../../../../components/manu/Nav";
import SideBar from "../../../../components/manu/Sidebar";
import EditPayment from "../../../../components/Edit/EditPayment";



const Editpayment: React.FC = () => {
    const paymentId = 'yourPaymentId';
   
return (
        <div className="flex">
            <SideBar />
            <div className="bg-black w-full">
                <Nav title="" />
                <div>
                    <div className=" navbar bg-black h-[76px] flex justify-between items-center px-4 py-2">
                        <h1 className="text-5xl text-green1 ml-14 mt-10 text-secondary">Edit Payment</h1>
                    </div>

                </div>
                <div className="flex flex-wrap justify-center">
                    <div className=" mt-5 w-[1500px] h-[1000px] rounded-3xl overflow-auto scrollable-div flex justify-center bg-sidebar backdrop-blur-sm">
                        <div className="flex flex-row items-start m-8 text-secondary ">
                             <EditPayment paymentId={paymentId} />
                        </div>
                    </div>
                </div>
            </div>
            
        </div>
    );
};

export default Editpayment;