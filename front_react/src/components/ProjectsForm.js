import { useState } from 'react';
import { Project } from '../App.js'

import React from 'react'
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import Autocomplete from '@material-ui/lab/Autocomplete';

export default function ProjectsForm(props) {
  const [index, setIndex] = useState(null);

  return (
    <div>
      <Autocomplete
        value={index}
        onChange={(event, newValue) => {
          setIndex(newValue);
        }}
        disablePortal
        id="autocomplete"
        options={props.projectsDoable.map((object, index) => index)} // the actual input value is the index in projects array
        getOptionLabel={(option) => props.projectsDoable[option].slug}
        sx={{ width: 300 }}
        renderInput={(params) => <TextField {...params} label="Project" />}
      />
      <Button variant="contained" color="primary"
        onClick={() => {
          if (index !== null) {
            if (props.projects === null || props.projects.length === 0)
              props.setProjects([new Project(props.projectsDoable[index].slug, props.projectsDoable[index].xp, 100, false), ]);
            else if (props.projects.findIndex(item => item.name === props.projectsDoable[index].slug) === -1)
                props.setProjects([...props.projects, new Project(props.projectsDoable[index].slug, props.projectsDoable[index].xp, 100, false)])
          }
        }}
      >
        Add Project
      </Button>
    </div>
  );
}
