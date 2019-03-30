import React, { Component } from 'react';
import './App.css';
import Table from 'react-bootstrap/Table';
import { wwdScore, waveScore, swpScore } from './functions';


const containerStyle = {
  margin: '10%',
  backgroundColor: '#FDFFFC',
  border: '1px solid lightgrey'
};

const conditionStyle = {
  backgroundColor: '#FDFFFC',
  border: '1px solid lightgrey'
};

const tableStyle = {
  marginLeft: '10%',
  marginRight: '10%'
};

class App extends Component {
  constructor(props){
    super(props);
    this.state ={
      surfData: [],
      totalScore: 0,
    }
  }

  componentDidMount() {
    fetch('http://localhost:3000', {
      // credentials: 'same-origin'
    })
      .then(response => response.json())
      .then(surfData => this.setState({ surfData }));

  }

  render() {
    
    return (
      <div className="App" >
        <div className="container" style={containerStyle} >
        <h1>Today's Surf Data</h1>
        <Table striped bordered hover size="sm" style={tableStyle} >
            <thead >
              <tr>
                <th>Date</th>
                <th>SwH</th>
                <th>SwP</th>
                <th>WWH</th>
                <th>WWD</th>
                <th>SwD</th>
                <th>WWP</th>
                <th>WAVE</th>
                <th>Score WWD</th>
                <th>Score Wave</th>
                <th>Score SwP</th>
                <th>Total Score</th>
              </tr>
            </thead>
          { this.state.surfData.map(surfDataRow => {
              if ((surfDataRow.Year == 2019) & (surfDataRow.Month == 3) & (surfDataRow.Day == 29)) {
                return (
                  <tbody>
                    <tr>
                      <td>{surfDataRow.Year}-{surfDataRow.Month}-{surfDataRow.Day} </td>
                      <td>{this.state.SwH = surfDataRow.SwH}</td>
                      <td>{surfDataRow.SwP}</td>
                      <td>{surfDataRow.WWH}</td>
                      <td>{surfDataRow.WWP}</td>
                      <td>{surfDataRow.SwD}</td>
                      <td>{surfDataRow.WWP}</td>
                      <td>{(surfDataRow.SwP * surfDataRow.SwH).toFixed(2)}</td>
                      <td>{ wwdScore(surfDataRow.WWD)}</td>
                      <td>{ waveScore(surfDataRow.SwP, surfDataRow.SwH)}</td>
                      <td>{ swpScore(surfDataRow.SwP)}</td>
                      <td>{ this.state.totalScore = ((swpScore(surfDataRow.SwP)+ wwdScore(surfDataRow.WWP) + waveScore(surfDataRow.SwP, surfDataRow.SwH))/3).toFixed(2) }</td>
                    </tr>
                  </tbody>
                )
              }
            })
          }
          </Table>
          <div className="conditions" style={conditionStyle}>
            <h3>Score of the Day</h3> 
            <p>
            { this.state.totalScore}
            </p>
          </div>
        </div>
      </div>
    );
  }
}

export default App;
