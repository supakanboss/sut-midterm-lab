import { Link as RouterLink } from "react-router-dom";
import Button from "@mui/material/Button";
import Grid from "@mui/material/Grid";
import Container from "@mui/material/Container";
import Paper from "@mui/material/Paper";
import Box from "@mui/material/Box";

import { Homebar } from "./Homebar";
import DataStudent from "./DataStudent";
import React from "react";
import { CssBaseline } from "@mui/material";

function Home() {
  return (
    <div className="Home" id="outer-container">
      <Homebar pageWrapId={"page-Home"} outerContainerId={"outer-container"} />
      <div id="page-Home">
        <React.Fragment>
          <CssBaseline />
          <Container maxWidth="lg">
            <Paper sx={{ padding: 2 }}>
              <Box display={"flex"}>
                <Box sx={{ flexGrow: 1 }}>
                  <Grid container spacing={2}>
                    <Grid item xs={2}>
                      <Button
                        variant="contained"
                        size="large"
                        fullWidth
                        component={RouterLink}
                        to="/DataStudent"
                      >
                        Datastudent
                      </Button>
                    </Grid>
                  </Grid>
                </Box>
              </Box>
            </Paper>
          </Container>
        </React.Fragment>
      </div>
    </div>
  );
}

export default Home;
