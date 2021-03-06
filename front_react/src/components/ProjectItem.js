import { useState } from 'react'

import { Project } from '../App';

import Typography from '@material-ui/core/Typography';
import Button from '@material-ui/core/Button';
import Box from '@material-ui/core/Box';
import Grid from '@material-ui/core/Grid';
import Slider from '@material-ui/core/Slider';
import Input from '@material-ui/core/Input';
import Checkbox from '@material-ui/core/Checkbox';
import FormControlLabel from '@material-ui/core/FormControlLabel';

export default function ProjectsItem(props) {
    const [percentage, setPercentage] = useState(props.projects[props.index].percentage);
    const [checked, setChecked] = useState(props.projects[props.index].checked);

    const updateProjects = (newValue) => {
        if (typeof newValue === 'boolean') {
            setChecked(newValue);
        }
        else {
            setPercentage(newValue);
        }
        let tmpProjects = props.projects;
        tmpProjects[props.index] = new Project(tmpProjects[props.index].name, tmpProjects[props.index].xp,
            typeof newValue !== 'boolean' ? newValue : tmpProjects[props.index].percentage,
            typeof newValue === 'boolean' ? newValue : tmpProjects[props.index].checked,
        );
        props.setProjects([...tmpProjects]);
    };

    const marks = [
        {
            value: 0,
            label: '0',
        },
        {
            value: 100,
            label: '100',
        },
        {
            value: 125,
            label: '125',
        },
    ];

    return (
        <Grid container spacing={3} alignItems="center" sx={{ margin: 2 }}>
            <Grid item>
                <Typography id="input-slider" >
                    {props.project.name}
                </Typography>
            </Grid>
            <Grid item>
                %
            </Grid>
            <Grid item xs>
                <Slider
                    aria-labelledby="input-slider"
                    value={percentage}
                    step={1}
                    min={0}
                    max={125}
                    color={percentage >= 100 ? "primary" : "secondary"}
                    onChange={(e, newPercentage) => {
                        updateProjects(newPercentage);
                    }}
                    marks={marks}
                />
            </Grid>
            <Grid item>
                <Input
                    value={percentage}
                    size="small"
                    onChange={(e) => {
                        if (e.target.value >= 0 && e.target.value <= 125) {
                            updateProjects(Number(e.target.value));
                        }
                        if (e.target.value > 125) {
                            updateProjects(125);
                        }
                    }}
                    inputProps={{
                        step: 1,
                        min: 0,
                        max: 125,
                        type: 'number',
                        'aria-labelledby': 'input-slider',
                    }}
                />
            </Grid>
            <Grid item>
                <FormControlLabel control={
                    <Checkbox color="primary"
                        value={checked}
                        onChange={(e) => { updateProjects(!checked) }} />}
                    label="Is my coa first ?"
                />
            </Grid>
            <Grid item>
                <Button variant="contained" color="primary"
                    onClick={(e) => {
                        props.setProjects(props.projects.filter(item => item !== props.project))
                    }}>
                    Remove
                </Button>
            </Grid>
        </Grid>
    );
}