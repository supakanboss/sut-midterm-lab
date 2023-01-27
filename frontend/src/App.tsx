import Home from "./components/Home";
import HomeStudent from "./components/Home-Student";
import HomeAdmin from "./components/Home-Admin";
import DataStudent from "./components/Student/DataStudent";
import CreateStudent from "./components/Student/CreateStudent";
import UpdateStudent from "./components/Student/UpdateStudent";
import SearchStudent from "./components/Student/SearchStudent";

import DataCourse from "./components/Course/DataCourse";
import CreateCourse from "./components/Course/CreateCourse";
import UpdateCourse from "./components/Course/UpdateCourse";
import SearchCourse from "./components/Course/SearchCourse";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

function App() {
  return (
        <Routes>
          
          <Route path="/" element={<Home />} />
          <Route path="/HomeStudent" element={<HomeStudent />} />
          <Route path="/HomeAdmin" element={<HomeAdmin />} />
          <Route path="/DataStudent" element={<DataStudent />} />
          <Route path="/CreateStudent" element={<CreateStudent />} />
          <Route path="/DataStudent/UpdateStudent/:id" element={<UpdateStudent />} />
          <Route path="/DataStudent/SearchStudent/:id" element={<SearchStudent />} />

          <Route path="/DataCourse" element={<DataCourse />} />
          <Route path="/CreateCourse" element={<CreateCourse />} />
          <Route path="/DataCourse/UpdateCourse/:id" element={<UpdateCourse />} />
          <Route path="/DataCourse/SearchCourse/:id" element={<SearchCourse />} />
          

        </Routes>

  );
}
export default App;
