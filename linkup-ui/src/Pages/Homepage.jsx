import React from 'react';
import {ButtonGroup, Button} from 'react-bootstrap';
import {Form, Col, Card} from 'react-bootstrap';
import axios from 'axios';
import Groups from './Groups';
import Events from './Events';

import MenuBar from '../Components/MenuBar';
export class HomePage extends React.Component {
  constructor(props) {
    super(props);
    this.mocking = true;
    this.data = this.mocking
      ? {
          terms: '',
          location: '',
          groupID: '',
          groupName: '',
          groupMembers: '',
          showGroups: '',
          showEvents: '',
        }
      : {};
    this.state = {
      response: '',
      groupCreateResponse: '',
      groupReadResponse: '',
      groupUpdateResponse: '',
      groupDeleteResponse: '',
    };
  }

  handleSubmit = () => {
    const terms = document.getElementById('termsId').value;
    const location = document.getElementById('locationId').value;
    const toSend = {
      terms: terms,
      location: location,
    };
    const outer = this;
    const instance = axios.create({timeout: 10000});
    console.log(toSend);
    instance.defaults.headers.common['Content-Type'] = 'application/json';
    instance
      .post('http://' + window.location.hostname + '/yelp', toSend)
      .then(res => {
        console.log(res);
        outer.setState({response: res});
      })
      .catch(err => {
        console.log(err);
        outer.setState({response: err});
      });
  };

  handleGroupButton = () => {
    this.setState({showGroups: <Groups />, showEvents: ''});
  };
  handleEventButton = () => {
    this.setState({showEvents: <Events />, showGroups: ''});
  };
  render() {
    const {response, showGroups, showEvents} = this.state;

    const yelp = (
      <Card>
        <Card.Body>
          <Card.Title>Access Yelp API</Card.Title>
          <Form>
            <Form.Group>
              <Col xs="5">
                <Form.Label>Enter Activity</Form.Label>
                <Form.Control type="text" id="termsId" />
              </Col>
            </Form.Group>

            <Form.Group>
              <Col xs="5">
                <Form.Label>Enter Location</Form.Label>
                <Form.Control type="text" id="locationId" />
              </Col>
            </Form.Group>
            <Col>
              <Button
                variant="success"
                type="button"
                onClick={e => this.handleSubmit(e)}
              >
                Submit
              </Button>
              <div>
                <br />
                {response.data}
              </div>
            </Col>
          </Form>
        </Card.Body>
      </Card>
    );

    return (
      <div>
        <MenuBar />
        {yelp}
        <Card>
          <ButtonGroup aria-label="Basic example" size="md">
            <Button variant="success" onClick={e => this.handleGroupButton(e)}>
              Groups
            </Button>
            <Button variant="success" onClick={e => this.handleEventButton(e)}>
              Events
            </Button>
          </ButtonGroup>
          {showGroups}
          {showEvents}
        </Card>
      </div>
    );
  }
}

export default HomePage;
