import { useState } from 'react';
import React from 'react'
import Box from '@material-ui/core/Box';
import Grid from '@material-ui/core/Grid';
import Slider from '@material-ui/core/Slider';
import Input from '@material-ui/core/Input';
import TextField from '@material-ui/core/TextField';
import Typography from '@material-ui/core/Typography';
import Checkbox from '@material-ui/core/Checkbox';
import FormControlLabel from '@material-ui/core/FormControlLabel';
import Button from '@material-ui/core/Button';
import Autocomplete from '@material-ui/lab/Autocomplete';

export default function Projects(props) {
  // let projects;
  // props.projectsDoable.forEach(function(item) {
  //     projects.push({ name: item.slug, xp: item.xp})
  //   });
  // console.log("projects");
  // console.log(projects);
  const [value, setValue] = useState(null);
  const [percentage, setPercentage] = React.useState(100);
  const [checked, setChecked] = React.useState(false);

  const handleSliderChange = (event, newValue) => {
    setPercentage(newValue);
  };

  const handleInputChange = (event) => {
    setPercentage(event.target.value === '' ? '' : Number(event.target.value));
  };

  const handleBlur = () => {
    if (percentage < 100) {
      setPercentage(100);
    } else if (percentage > 125) {
      setPercentage(125);
    }
  };

  const handleCheckBox = (event) => {
    setChecked(event.target.checked);
  };

  return (
    <div>
      <Autocomplete

        // I think we can do this once we implement the button to count the project
        // onChange={(e, slug) => {for (var i = 0; i < props.projectsDoable.length; i++) {
        //   if (slug === props.projectsDoable[i].slug)
        //       props.projectsDone.push({"name":slug, "xp":props.projectsDone[i].xp});
        // }}}

        value={value}
        onChange={(event, newValue) => {
          setValue(newValue);
        }}
        disablePortal
        id="autocomplete"
        options={props.projectsDoable.map((object, index) => index)} // the actual input value is the index in projects array
        getOptionLabel={(option) => props.projectsDoable[option].slug}
        sx={{ width: 300 }}
        renderInput={(params) => <TextField {...params} label="Project" />}
      />
      <Box sx={{ width: 250 }}>
        <Typography id="input-slider" gutterBottom>
          Validation Percentage
        </Typography>
        <Grid container spacing={2} alignItems="center">
          <Grid item>
            %
          </Grid>
          <Grid item xs>
            <Slider
              value={typeof percentage === 'number' ? percentage : 0}
              onChange={handleSliderChange}
              aria-labelledby="input-slider"
              defaultValue={100}
              step={1}
              min={100}
              max={125}
            />
          </Grid>
          <Grid item>
            <Input
              value={percentage}
              size="small"
              onChange={handleInputChange}
              onBlur={handleBlur}
              inputProps={{
                step: 1,
                min: 100,
                max: 125,
                type: 'number',
                'aria-labelledby': 'input-slider',
              }}
            />
          </Grid>
        </Grid>
      </Box>
      <FormControlLabel control={<Checkbox color="primary" onChange={handleCheckBox} />} label="Is my coa first ?" />
      <Button variant="contained" color="primary"
        onClick={() => {
          // handleClick(); // -> addProject
        }}
      >
        Submit Project
      </Button>
    </div>
  );
}
