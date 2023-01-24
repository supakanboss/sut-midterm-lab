import Home from "./component/Home";
import Homebar from "./component/Homebar";
import DataStudent from "./component/DataStudent";
import CreateStudent from "./component/CreateStudent";
import UpdateStudent from "./component/UpdateStudent";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

function App() {
  return (
        <Routes>
          
          <Route path="/" element={<Home />} />
          <Route path="/DataStudent" element={<DataStudent />} />
          <Route path="/CreateStudent" element={<CreateStudent />} />
          <Route path="/UpdateStudent" element={<UpdateStudent />} />

        </Routes>

  );
}
export default App;
