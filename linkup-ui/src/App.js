import React, {Component} from 'react';
import axios from 'axios';

class App extends Component {
  constructor(props) {
    super(props);
    this.state = {
      response: '',
    };
  }

  //Will use something like this once we actually start implementing webpages with 'submit' buttons to press
  // handleSubmit = () => {
  //   const toSend = {
  //     terms: "restaurant",
  //     location: "Boston"
  //   };
  //   const instance = axios.create({timeout: 10000});
  //   instance.defaults.headers.common['Content-Type'] = 'application/json';
  //   instance.post('/yelp', toSend).
  //   then((res) => {
  //     this.setState({response: res});
  //   }).catch((err) => {
  //     this.setState({response: err});
  //   });
  // }
  componentDidMount() {
    const toSend = {
      terms: 'restaurant',
      location: 'Boston',
    };
    const instance = axios.create({timeout: 10000});
    instance.defaults.headers.common['Content-Type'] = 'application/json';
    instance
      .post('http://' + window.location.hostname + '/yelp', toSend)
      .then(res => {
        this.setState({response: res});
      })
      .catch(err => {
        this.setState({response: err});
      });
  }
  render() {
    const {response} = this.state;
    return (
      <div className="App">
        <h1>Making Yelp Post Request</h1>
        {response.data}
      </div>
    );
  }
}

export default App;
