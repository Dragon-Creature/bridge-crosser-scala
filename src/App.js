import "./App.css";
import React, { useState } from "react";
import uuid from "react-uuid";

function App() {
  const [errorMessage, setErrorMessage] = useState("");
  const [bridges, setBridges] = useState([{ id: uuid(), hikers: [] }]);
  const [result, setResult] = useState(undefined);

  const createBridge = () => {
    let newBridges = JSON.parse(JSON.stringify(bridges));
    newBridges.push({ id: uuid(), length_in_feet: 0, hikers: [] });
    setBridges(newBridges);
  };

  const createHiker = (bridge, idx) => {
    bridge.hikers.push({ id: uuid() });
    bridges[idx] = bridge;
    let newBridges = JSON.parse(JSON.stringify(bridges));
    setBridges(newBridges);
  };

  const updateLengthInFeet = (value, idx) => {
    let newBridges = JSON.parse(JSON.stringify(bridges));
    newBridges[idx].length_in_feet = Number(value);
    setBridges(newBridges);
  };

  const updateSpeedInFeet = (value, idx, idy) => {
    let newBridges = JSON.parse(JSON.stringify(bridges));
    newBridges[idx].hikers[idy].speed_feet_in_minutes = Number(value);
    setBridges(newBridges);
  };

  const Submit = () => {
    setErrorMessage("");
    setResult(undefined);
    fetch("/api/calculate_crossing", {
      method: "post",
      body: JSON.stringify({ bridges }),
    })
      .then((response) => {
        if (response.ok) {
          response.json().then((response) => setResult(response));
        } else {
          response
            .json()
            .then((response) =>
              setErrorMessage(
                <p className={"Error-Message"}>{response.message}</p>
              )
            );
        }
      })
      .catch((err) =>
        setErrorMessage(
          <p className={"Error-Message"}>
            unexpected error: unable to connect to backend
          </p>
        )
      );
  };

  let resultRender = <div></div>;
  if (result !== undefined) {
    resultRender = (
      <div>
        <p>Total travel time</p>
        <p>{result.total_travel_time}</p>
        {result.bridge_results.map((bridge) => (
          <div>
            <p>Total travel time</p>
            <p>{bridge.total_travel_time}</p>
            <p>Number of hikers crossed</p>
            <p>{bridge.number_of_hikers_crossed}</p>
            <p>Id of hikers</p>
            <p>{bridge.id_of_hikers}</p>
            <p>Length in feet</p>
            <p>{bridge.length_in_feet}</p>
          </div>
        ))}
      </div>
    );
  }
  return (
    <div className="App">
      {errorMessage}
      {resultRender}
      <div className={"Bridges"}>
        {bridges.map((bridge, idx) => (
          <div className={"Bridge-Container"} key={bridge.id}>
            <p>Bridge {idx + 1}</p>
            <div className={"Bridge"}>
              <p>Length in feet</p>
              <input
                type={"text"}
                value={bridge.length_in_feet}
                onChange={(e) => {
                  updateLengthInFeet(e.target.value, idx);
                }}
              />
              <button
                data-testid="add-hiker"
                onClick={() => createHiker(bridge, idx)}
              >
                Add Hiker
              </button>
              {bridge.hikers.map((hiker, idy) => (
                <div className={"Hiker-Container"} key={hiker.id}>
                  <p>Hiker {idy + 1}</p>
                  <div className={"Hiker"}>
                    <p>ID</p>
                    <input type={"text"} value={hiker.id} readOnly={true} />
                    <p>Speed feet in minutes</p>
                    <input
                      type={"text"}
                      value={hiker.speed_feet_in_minutes}
                      onChange={(e) => {
                        updateSpeedInFeet(e.target.value, idx, idy);
                      }}
                    />
                  </div>
                </div>
              ))}
            </div>
          </div>
        ))}
      </div>
      <div className={"Control-Buttons"}>
        <button onClick={createBridge} data-testid="add-bridge">
          Add Bridge
        </button>
        <button onClick={Submit}>Submit</button>
      </div>
    </div>
  );
}

export default App;
