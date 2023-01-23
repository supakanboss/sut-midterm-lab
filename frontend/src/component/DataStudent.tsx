import React, { useEffect, useState } from "react";
import CssBaseline from "@mui/material/CssBaseline";
import Box from "@mui/material/Box";
import Container from "@mui/material/Container";
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
import { Link as RouterLink } from "react-router-dom";

import { StudentInterface } from "../model/IStudent";

import Homebar from "./Homebar";
import CreateStudent from "./CreateStudent";

function DataStudent() {
  /////////////////////////////////////////////////////

  const [Studentstable, setStudentstable] = useState<StudentInterface[]>([]);

  /////////////////////////////////////////////////////
  const apiUrl = "http://localhost:8080";
  const requestOpionsGet = {
    method: "GET",
    headers: { "Content-Type": "application/json" },
  };
  /////////////////////////////////////////////////////

  //แสดงข้อมูล student ทั้งหมด
  const feachStudentstable = async () => {
    fetch(`${apiUrl}/student_table`, requestOpionsGet)
      .then((response) => response.json())
      .then((result) => {
        console.log(result.data);
        setStudentstable(result.data);
      });
  };

  /////////////////////////////////////////////////////



  /////////////////////////////////////////////////////

  useEffect(() => {
    feachStudentstable();
  }, []);

  /////////////////////////////////////////////////////

  return (
    <div>
      <Homebar />
      <React.Fragment>
        <CssBaseline />
        <Container maxWidth="lg" sx={{ padding: 2 }}>
          <Paper sx={{ padding: 2 }}>
            <Box display={"flex"}>
              <Box sx={{ flexGrow: 1 }}>
                <Typography variant="h5" gutterBottom>
                  Student
                </Typography>
              </Box>
              <Box>
                <Button
                  variant="contained"
                  component={RouterLink}
                  to="/CreateStudent"
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
                  <TableCell align="center">Student_Number</TableCell>
                  <TableCell align="center">Student_Name</TableCell>
                  <TableCell align="center">Student_Identity_Card</TableCell>
                  <TableCell align="center">Student_Tel</TableCell>
                  <TableCell align="center"></TableCell>
                </TableRow>
              </TableHead>
              <TableBody>
                {Studentstable.map((row) => (
                  <TableRow
                    key={row.ID}
                    sx={{ "&:last-child td, &:last-child th": { border: 0 } }}
                  >
                    <TableCell align="center">{row.ID}</TableCell>
                    <TableCell align="center">{row.Student_Number}</TableCell>
                    <TableCell align="center">{row.Student_Name}</TableCell>
                    <TableCell align="center">
                      {row.Student_Identity_Card}{" "}
                    </TableCell>
                    <TableCell align="center">{row.Student_Tel} </TableCell>
                    <TableCell align="center">
                      <ButtonGroup>
                        <Button>update</Button>
                        <Button>
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
  );
}

export default DataStudent;
