import React from 'react';
import {Button} from 'react-bootstrap';
import axios from 'axios';
export class HomePage extends React.Component {
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
    console.log(response.data);
    let showButton = '';
    if (response.data !== undefined) {
      showButton = (
        <div>
          <br />
          <Button>Post Successful</Button>
        </div>
      );
    }
    return (
      <div>
        <div>
          <h1>Making Yelp Post Request</h1>
          {response.data}
        </div>
        <div>{showButton}</div>
      </div>
    );
  }
}

export default HomePage;
