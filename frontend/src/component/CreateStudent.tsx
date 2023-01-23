import React, { useEffect, useState } from "react";
import CssBaseline from "@mui/material/CssBaseline";
import Box from "@mui/material/Box";
import Container from "@mui/material/Container";
import Typography from "@mui/material/Typography";
import Button from "@mui/material/Button";
import Paper from "@mui/material/Paper";
import Grid from "@mui/material/Grid";
import TextField from "@mui/material/TextField";

import Homebar from "./Homebar";

function CreateStudent() {
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
                  CreateStudent
                </Typography>
              </Box>
            </Box>
          </Paper>
        </Container>
        <Container maxWidth="lg">
          <Paper sx={{ padding: 2 }}>
            <Box display={"flex"}>
              <Box sx={{ flexGrow: 1 }}>
                <Grid container spacing={2}>
                  <Grid item xs={12}>
                    <h4>ประวัติส่วนตัวนักศึกษา</h4>
                    <hr />
                  </Grid>
                  <Grid item xs={2}>
                    <p>Prefix </p>
                    <TextField
                      fullWidth
                      id="outlined-basic"
                      label="Prefix"
                      variant="outlined"
                      name="Prefix"
                      //value={}
                      //onChange={}
                    />
                  </Grid>
                  <Grid item xs={4}>
                    <p>ชื่อ-สกุล</p>
                    <TextField
                      fullWidth
                      id="Member_Name"
                      type="string"
                      label="ชื่อ-สกุล"
                      variant="outlined"
                      name="Member_Name"
                      //value={}
                      //onChange={}
                    />
                  </Grid>
                  <Grid item xs={6}></Grid>
                  <Grid item xs={2}>
                    <p>Gender </p>
                    <TextField
                      fullWidth
                      id="outlined-basic"
                      label="Genden"
                      variant="outlined"
                    />
                  </Grid>
                  <Grid item xs={4}>
                    <p>Birthday</p>
                    <TextField
                      fullWidth
                      id="Member_Name"
                      type="date"
                      label=""
                      variant="outlined"
                      name="Member_Name"
                      //value={}
                      //onChange={}
                    />
                  </Grid>
                  <Grid item xs={6}></Grid>
                  <Grid item xs={6}>
                    <p>tel</p>
                    <TextField
                      fullWidth
                      id="Member_Name"
                      type="string"
                      label="tel"
                      variant="outlined"
                      name="Member_Name"
                      //value={}
                      //onChange={}
                    />
                  </Grid>
                  <Grid item xs={6}></Grid>
                  <Grid item xs={6}>
                    <p>identity card</p>
                    <TextField
                      fullWidth
                      id="Member_Name"
                      type="string"
                      label="รหัสบัตรประชาชน"
                      variant="outlined"
                      name="Member_Name"
                      //value={}
                      //onChange={}
                    />
                  </Grid>
                  <Grid item xs={6}></Grid>
                  <Grid item xs={6}>
                    <p>father</p>
                    <TextField
                      fullWidth
                      id="Member_Name"
                      type="string"
                      label="ชื่อ-สกุล บิดา"
                      variant="outlined"
                      name="Member_Name"
                      //value={}
                      //onChange={}
                    />
                  </Grid>
                  <Grid item xs={6}></Grid>
                  <Grid item xs={6}>
                    <p>Mother</p>
                    <TextField
                      fullWidth
                      id="Member_Name"
                      type="string"
                      label="ชื่อ-สกุล มารดา"
                      variant="outlined"
                      name="Member_Name"
                      //value={}
                      //onChange={}
                    />
                  </Grid>
                  <Grid item xs={6}></Grid>
                  <Grid item xs={3}>
                    <p>nationality</p>
                    <TextField
                      fullWidth
                      id="Member_Name"
                      type="string"
                      label="สัญชาติ"
                      variant="outlined"
                      name="Member_Name"
                      //value={}
                      //onChange={}
                    />
                  </Grid>
                  <Grid item xs={3}>
                    <p>religion</p>
                    <TextField
                      fullWidth
                      id="Member_Name"
                      type="string"
                      label="ศาสนา"
                      variant="outlined"
                      name="Member_Name"
                      //value={}
                      //onChange={}
                    />
                  </Grid>
                  <Grid item xs={6}></Grid>
                  <Grid item xs={6}>
                    <p>province</p>
                    <TextField
                      fullWidth
                      id="Member_Name"
                      type="string"
                      label="จังหวัด"
                      variant="outlined"
                      name="Member_Name"
                      //value={}
                      //onChange={}
                    />
                  </Grid>
                  <Grid item xs={6}></Grid>
                  <Grid item xs={6}>
                    <p>address</p>
                    <TextField
                      fullWidth
                      id="Member_Name"
                      type="string"
                      label="ที่อยู่"
                      variant="outlined"
                      name="Member_Name"
                      //value={}
                      //onChange={}
                    />
                  </Grid>
                  <Grid item xs={6}></Grid>
                  <Grid item xs={12}>
                    <hr />
                    <h4>ข้อมูลนักศึกษา</h4>
                    
                  </Grid>
                  <Grid item xs={3}>
                    <p>institute</p>
                    <TextField
                      fullWidth
                      id="Member_Name"
                      type="string"
                      label="สำนักวิชา"
                      variant="outlined"
                      name="Member_Name"
                      //value={}
                      //onChange={}
                    />
                  </Grid>
                  <Grid item xs={3}>
                    <p>branch</p>
                    <TextField
                      fullWidth
                      id="Member_Name"
                      type="string"
                      label="สาขาวิชา"
                      variant="outlined"
                      name="Member_Name"
                      //value={}
                      //onChange={}
                    />
                  </Grid>
                  <Grid item xs={6}></Grid>
                  <Grid item xs={3}>
                    <p>course</p>
                    <TextField
                      fullWidth
                      id="Member_Name"
                      type="string"
                      label="หลักสูตร"
                      variant="outlined"
                      name="Member_Name"
                      //value={}
                      //onChange={}
                    />
                  </Grid>
                  <Grid item xs={3}>
                    <p>degree</p>
                    <TextField
                      fullWidth
                      id="Member_Name"
                      type="string"
                      label="ระดับการศึกษา"
                      variant="outlined"
                      name="Member_Name"
                      //value={}
                      //onChange={}
                    />
                  </Grid>
                  <Grid item xs={6}></Grid>
                  <Grid item xs={6}>
                    <p>Year of Entry</p>
                    <TextField
                      fullWidth
                      id="Member_Name"
                      type="date"
                      label=""
                      variant="outlined"
                      name="Member_Name"
                      //value={}
                      //onChange={}
                    />
                  </Grid>
                  <Grid item xs={6}></Grid>
                  <Grid item xs={6}>
                    <p>student Number</p>
                    <TextField
                      fullWidth
                      id="Member_Name"
                      type="string"
                      label="รหัสนักศึกษา"
                      variant="outlined"
                      name="Member_Name"
                      //value={}
                      //onChange={}
                    />
                  </Grid>
                  <Grid item xs={6}></Grid>
                  <Grid item xs={3}>
                  <Button variant="contained" size="large" fullWidth>submit</Button>
                  </Grid>
                  <Grid item xs={3}>
                  <Button variant="contained" size="large" fullWidth>back</Button>
                  </Grid>
                  <Grid item xs={6}></Grid>
                </Grid>
              </Box>
            </Box>
          </Paper>
        </Container>
      </React.Fragment>
    </div>
  );
}
export default CreateStudent;
