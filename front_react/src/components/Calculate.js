export default function Calculate(props) {
    let level = props.level;
    let eplvl = [
        0,
        462,
        2688,
        5885,
        11777,
        29217,
        46255,
        63559,
        74340,
        85483,
        95000,
        105630,
        124446,
        145782,
        169932,
        197316,
        228354,
        263508,
        303366,
        348516,
        399672,
        457632,
        523320,
        597786,
        682164,
        777756,
        886074,
        1008798,
        1147902,
        1305486,
        1484070,
        -1
    ];
    let lvlstart_int = parseInt(String(level));
    let fract_part = level - lvlstart_int;
    let xp_total = eplvl[lvlstart_int] +
    ((eplvl[lvlstart_int + 1] - eplvl[lvlstart_int]) * fract_part);;
    if (props.projects !== null && props.projects.length > 0) {
        for (let i = 0; i < props.projects.length; i++) {
            if (props.projects[i].percentage >= 100) {
                lvlstart_int = parseInt(String(level));
                fract_part = level - lvlstart_int;
                let coabonus = props.projects[i].checked ? 1.042 : 1;
                xp_total = props.projects[i].xp * (props.projects[i].percentage / 100) * coabonus
                    + eplvl[lvlstart_int] +
                    ((eplvl[lvlstart_int + 1] - eplvl[lvlstart_int]) * fract_part);
                let j = eplvl.findIndex(exp => exp > xp_total) - 1;
                level = j + ((xp_total - eplvl[j]) / (eplvl[j + 1] - eplvl[j]));
            }
        }
    }
    return (
            <div>
                <div>Starting level : {props.level}</div>
                <div>Ending level : {level.toFixed(2)}</div>
                <div>You still have {parseInt(String(eplvl[21] - xp_total))} xp to go until level 21</div>
            </div>
    );
}
