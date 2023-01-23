import { Link as RouterLink } from "react-router-dom";
import Button from "@mui/material/Button";
import Grid from "@mui/material/Grid";
import Container from "@mui/material/Container";
import Paper from "@mui/material/Paper";
import Box from "@mui/material/Box";

import Homebar from "./Homebar";
import DataStudent from "./DataStudent";

function Home() {
  return (
    <div>
      <Homebar />
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
    </div>
  );
}

export default Home;
