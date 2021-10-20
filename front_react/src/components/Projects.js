import TextField from '@material-ui/core/TextField';
import Autocomplete from '@material-ui/lab/Autocomplete'
import { useState } from 'react';

export default function Projects(props) {
    // let projects;
    // props.projectsDoable.forEach(function(item) {
    //     projects.push({ name: item.slug, xp: item.xp})
    //   });
    // console.log("projects");
    // console.log(projects);
    const [value, setValue] = useState(null);

    return (
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
      );
}
