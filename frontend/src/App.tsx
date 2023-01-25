import Home from "./component/Home";
import Homebar from "./component/Homebar";
import DataStudent from "./component/DataStudent";
import CreateStudent from "./component/CreateStudent";
import UpdateStudent from "./component/UpdateStudent";
import SearchStudent from "./component/SearchStudent";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

function App() {
  return (
        <Routes>
          
          <Route path="/" element={<Home />} />
          <Route path="/DataStudent" element={<DataStudent />} />
          <Route path="/CreateStudent" element={<CreateStudent />} />
          <Route path="/DataStudent/UpdateStudent/:id" element={<UpdateStudent />} />
          <Route path="/DataStudent/SearchStudent/:id" element={<SearchStudent />} />

        </Routes>

  );
}
export default App;
