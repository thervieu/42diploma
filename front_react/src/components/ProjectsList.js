import ProjectItem from './ProjectItem'

import List from '@material-ui/core/List';

export default function ProjectsList(props) {
    return (
        <List >
            {
                props.projects.map(item => {
                    return (
                        <ProjectItem project={item} projects={props.projects} setProjects={props.setProjects} />
                    );
                })
            }
        </List>
    );
}
