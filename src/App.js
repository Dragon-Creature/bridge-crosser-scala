import './App.css';
import React, {useState} from "react";
import uuid from 'react-uuid'

function App() {
    const [bridges, setBridges] = useState([{ id: uuid(), hikers: []}]);

    const createBridge = () => {
      let newBridges = JSON.parse(JSON.stringify(bridges));
        newBridges.push({id: uuid(), length_in_feet: 0, hikers: [{id: uuid()}]});
      setBridges(newBridges);
    }

    const createHiker = (bridge, idx) => {
        bridge.hikers.push({id: uuid()});
        bridges[idx] = bridge;
        let newBridges = JSON.parse(JSON.stringify(bridges));
        setBridges(newBridges);
    }

    const updateLengthInFeet = (value, idx) => {
        let newBridges = JSON.parse(JSON.stringify(bridges));
        newBridges[idx].length_in_feet = value;
        setBridges(newBridges);
    }

    const updateSpeedInFeet = (value, idx, idy) => {
        let newBridges = JSON.parse(JSON.stringify(bridges));
        newBridges[idx].hikers[idy].speed_feet_in_minutes = value;
        setBridges(newBridges);
    }

    const Submit = () => {
        fetch('/api/calculate_crossing', {
            method: 'post',
            body: JSON.stringify(bridges)
        })
    }

  return (
    <div className="App">
        {bridges.map((bridge, idx) => <div key={bridge.id}>
            <p>Length in feet</p>
            <input type={"text"} value={bridge.length_in_feet} onChange={(e) => {
                updateLengthInFeet(e.target.value, idx);
            }
            }/>
            <button onClick={() => createHiker(bridge, idx)}>Add Hiker</button>
            {bridge.hikers.map((hiker, idy) =>
                <div>
                    <p>ID</p>
                    <input type={"text"} value={hiker.id}/>
                    <p>Speed feet in minutes</p>
                    <input type={"text"} value={hiker.speed_feet_in_minutes} onChange={(e) => {
                        updateSpeedInFeet(e.target.value, idx, idy);
                    }
                    }/>
                </div>
            )}
        </div>)}

        <button onClick={createBridge}>Add Bridge</button>
        <button onClick={Submit}>Submit</button>
    </div>
  );
}

export default App;
