import TextField from '@material-ui/core/TextField';
import Autocomplete from '@material-ui/lab/Autocomplete'
import { useState } from 'react'

export default function Projects(props) {
    // let projects;
    // props.projectsDoable.forEach(function(item) {
    //     projects.push({ name: item.slug, xp: item.xp})
    //   });
    // console.log("projects");
    // console.log(projects);
    const [inputValue, setinputValue] = useState("");
    return (
        <Autocomplete
          inputValue={inputValue}
          onInputChange={(e) => {if (e && e.target !== null) setinputValue(e.target.value)}}
          onChange={(e, slug) => {for (var i = 0; i < props.projectsDoable.length; i++) {
            if (slug === props.projectsDoable[i].slug)
                props.projectsDone.push({"name":slug, "xp":props.projectsDone[i].xp});
          }}}
          id="autocomplete"
          options={props.projectsDoable}
          getOptionLabel={(option) => option.slug}
          sx={{ width: 300 }}
          renderInput={(params) => <TextField {...params} label="Project" />}
          open={inputValue.length > 0}
        />
      );
}
