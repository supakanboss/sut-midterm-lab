import Home from "./component/Home";
import CreateStudent from "./component/CreateStudent";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

function App() {
  return (
        <Routes>
          
          <Route path="/" element={<Home />} />
          <Route path="/CreateStudent" element={<CreateStudent />} />

        </Routes>

  );
}
export default App;
