import React, { Component } from 'react';
import axios from 'axios';

class App extends Component {

  constructor(props){
    super (props)
    this.state = {
      response: ''
    }
  }
  // handleSubmit = () => { POST REQUEST
  //   const toSend = {
  //     terms: "restaurant",
  //     location: "Boston"
  //   }
  //   const instance = axios.create({timeout: 10000});
  //   instance.defaults.headers.common['Content-Type'] = 'application/json';
  //   instance.post('/yelp', toSend).
  //   then((response) => {
  //     console.log('Success \n' + response)
  //   }).catch((error) => {
  //     console.log('Error \n' + error);
  //   });
  // }

  componentDidMount() {
    axios.get('http://' + window.location.hostname + '/yelp?terms=restaurant&location=Boston').
    then(res => {
      this.setState({response: res});
      console.log(res);
    }).catch(err => {
      this.setState({response: err})
      console.log(err);
    });
  }


  render() {
    const { response } = this.state;
    console.log('This is the response ' + response);
    return (
      <div className="App">
        <h1>Making Yelp Post Request</h1>

      {response.data}

      </div>
    );
  }
}

export default App;