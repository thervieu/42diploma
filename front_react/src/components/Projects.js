export default function Projects(props) {

    return (
    <div>
        <table>
            {props.projectsDoable.map(item => {
                return (
                    <td>{item.slug}</td>
                )
            })}
        </table>
    </div> );
}