import React, { useEffect, useState } from "react";
import CssBaseline from "@mui/material/CssBaseline";
import Box from "@mui/material/Box";
import Container from "@mui/material/Container";
import TextField from "@mui/material/TextField";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Paper from "@mui/material/Paper";
import Table from "@mui/material/Table";
import TableBody from "@mui/material/TableBody";
import TableCell from "@mui/material/TableCell";
import TableContainer from "@mui/material/TableContainer";
import TableHead from "@mui/material/TableHead";
import TableRow from "@mui/material/TableRow";
import DeleteOutlineIcon from "@mui/icons-material/DeleteOutline";
import { ButtonGroup } from "@mui/material";
import { Link as RouterLink, useNavigate } from "react-router-dom";
import SearchIcon from "@mui/icons-material/Search";
import { ThemeProvider, createTheme } from "@mui/material/styles";
import { HiHome } from "react-icons/hi";
import { BiSearchAlt } from "react-icons/bi";


import { CourseInterface } from "../../models/ICourse";

import { Adminbar } from "../Bar-Admin";

const Theme = createTheme({
  palette: {
    primary: {
      main: "#313131",
    },
    secondary: {
      main: "#C70039",
    },
    info: {
      main: "#164DC9",
    },
  },
});

function DataCourse() {
  /////////////////////////////////////////////////////

  let navigate = useNavigate();

  const [coursetable, setCoursetable] = useState<CourseInterface[]>([]);
  const [course, setCourse] = useState<Partial<CourseInterface>>({});

  /////////////////////////////////////////////////////
  const apiUrl = "http://localhost:8080";
  const requestOptionsGet = {
    method: "GET",
    headers: { "Content-Type": "application/json" },
  };
  /////////////////////////////////////////////////////

  //แสดงข้อมูล student ทั้งหมด
  const feachCoursetable = async () => {
    fetch(`${apiUrl}/course_table`, requestOptionsGet)
      .then((response) => response.json())
      .then((result) => {
        console.log(result.data);
        setCoursetable(result.data);
      });
  };

  const DeleteCourse = (id: string) => {
    console.log(id);
    const requestOptions = {
      method: "DELETE",
      headers: { "Content-Type": "application/json" },
    };

    fetch(`${apiUrl}/delete_course/${id}`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        console.log(res);
        if (res.data) {
          // setSuccess(true);
          alert(`Are you sure delete id: ${id}`);
          setTimeout(() => {
            window.location.href = "/DataCourse";
          }, 500);
        } else {
          // setError(true);
        }
      });
  };

  /////////////////////////////////////////////////////

  const handleInputChange = (
    event: React.ChangeEvent<{ id?: string; value: any }>
  ) => {
    const id = event.target.id as keyof typeof course;
    const { value } = event.target;
    setCourse({ ...course, [id]: value });
  };

  /////////////////////////////////////////////////////

  useEffect(() => {
    feachCoursetable();
  }, []);

  /////////////////////////////////////////////////////

  return (
    <div className="DataCourse" id="outer-container">
      <ThemeProvider theme={Theme}>
        <Adminbar
          pageWrapId={"page-DataCourse"}
          outerContainerId={"outer-container"}
        />
        <div id="page-DataCourse">
          <React.Fragment>
            <CssBaseline />
            <Container maxWidth="lg">
              <Paper sx={{ padding: 1 }}>
                <Box display={"flex"}>
                  <Box>
                    <Typography variant="h4" gutterBottom >
                    <Button 
                    color="inherit"
                    component={RouterLink}
                    to="/HomeAdmin"
                    >
                      <HiHome size="30"/ >
                    </Button>
                      Course
                    </Typography>
                  </Box>
                  <Box sx={{marginLeft:25}}>
                    <Typography variant="h4" gutterBottom >
                    <TextField
                        fullWidth
                        id="Course_Name"
                        type="string"
                        label="ค้นหา หลักสูตร"
                        variant="standard"
                        name="Course_Name"
                        value={course.Course_Name}
                        onChange={handleInputChange}
                      />
                    </Typography>
                  </Box>
                    <Button color="inherit">
                      <BiSearchAlt size="30"/ >
                    </Button>
                  <Box sx={{marginLeft:44}}>
                    <Button
                      variant="contained"
                      component={RouterLink}
                      to="/CreateCourse"
                      color="primary"
                      size="large"
                    >
                      create
                    </Button>
                  </Box>
                </Box>
              </Paper>
              <TableContainer component={Paper}>
                <Table sx={{ minWidth: 650 }} aria-label="simple table">
                  <TableHead>
                    <TableRow>
                      <TableCell align="center">ID</TableCell>
                      <TableCell align="center">Course_Name</TableCell>
                      <TableCell align="center">Course_Degree</TableCell>
                      <TableCell align="center">Course_Year</TableCell>
                      <TableCell align="center">Course_Credit</TableCell>
                      <TableCell align="center">Course_Teacher</TableCell>
                      <TableCell align="center"></TableCell>
                    </TableRow>
                  </TableHead>
                  <TableBody>
                    {coursetable.map((row) => (
                      <TableRow
                        key={row.ID}
                        sx={{
                          "&:last-child td, &:last-child th": { border: 0 },
                        }}
                      >
                        <TableCell align="center">{row.ID}</TableCell>
                        <TableCell align="center">{row.Course_Name}</TableCell>
                        <TableCell align="center">{row.Degree.Degree_Name}</TableCell>
                        <TableCell align="center">{row.Course_Year}</TableCell>
                        <TableCell align="center">{row.Course_Credit}</TableCell>
                        <TableCell align="center">{row.Course_Teacher} </TableCell>
                        <TableCell align="center">
                          <ButtonGroup>
                            <Button
                              onClick={() =>
                                navigate(`UpdateCourse/${row.ID}`)
                              }
                              color="info"
                            >
                              update
                            </Button>
                            <Button
                              onClick={() =>
                                navigate(`SearchCourse/${row.ID}`)
                              }
                            >
                              <SearchIcon />
                            </Button>
                            <Button onClick={() => DeleteCourse(row.ID + "")} color="secondary">
                              <DeleteOutlineIcon />
                            </Button>
                          </ButtonGroup>
                        </TableCell>
                      </TableRow>
                    ))}
                  </TableBody>
                </Table>
              </TableContainer>
            </Container>
          </React.Fragment>
        </div>
      </ThemeProvider>
    </div>
  );
}

export default DataCourse;
