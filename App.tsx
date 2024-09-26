import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import Home from "./pages/User/Home";
import Login from "./pages/User/Login";
import Package from "./pages/User/Package"; 
import Payment from "./pages/User/Payment"; 
import CreatePackage from "./pages/Admin/package/create";
import EditPackage from "./pages/Admin/package/edit";
import EditPayment from "./pages/Admin/payment/edit";


const App: React.FC = (): JSX.Element => {
    return (
        <Router>
            <Routes>
                {/* Correct path to /payment */}
                <Route path="/" element={<Home />} />
                <Route path="/home" element={<Home />} />
                <Route path="/login" element={<Login />} />
                <Route path="/package" element={<Package />} />
                <Route path="/payment" element={<Payment />} />
                <Route path="/package/create" element={<CreatePackage/>}/>
                <Route path="/package/edit" element={<EditPackage/>}/>
                <Route path="/payment/edit" element={<EditPayment/>}/>
            </Routes>
        </Router>
    );
}

export default App;
