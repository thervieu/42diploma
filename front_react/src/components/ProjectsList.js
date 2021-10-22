import ProjectItem from './ProjectItem'

import List from '@material-ui/core/List';

export default function ProjectsList(props) {
    return (
        props.projects !== null && props.projects.length > 0 ?
            <List >
                {
                    props.projects.map((item, index) => {
                        return (
                            <ProjectItem index={index} project={item} projects={props.projects} setProjects={props.setProjects} />
                        );
                    })
                }
            </List>
            :
            <div>Please choose a project</div>
    );
}
