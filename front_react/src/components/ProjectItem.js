import Button from '@material-ui/core/Button';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';

export default function ProjectsItem(props) {
    return (
        <ListItem >
            <ListItemText primary={props.project.name} />
            <ListItemText primary={"XP : " + String(props.project.xp)} />
            <ListItemText primary={"Percentage : " + String(props.project.percentage)} />
            <ListItemText primary={props.project.checked ? "Coalition first" : "Coa not first"} />
            <Button variant="contained" color="primary"
                onClick={() => {
                    props.setProjects(props.projects.filter(item => item !== props.project))
                }}>Remove</Button>
        </ListItem>
    );
}