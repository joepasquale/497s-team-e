import React from 'react';
import {Button} from 'react-bootstrap';
import {Form, Col, Card} from 'react-bootstrap';
import axios from 'axios';

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
        }
      : {};
    this.state = {
      response: '',
      groupResponse: '',
    };
  }
  handleSubmitGroup = () => {
    const id = document.getElementById('groupID').value;
    const name = document.getElementById('groupName').value;
    const members = document.getElementById('groupMembers').value;

    const groupMembers = members.split(',').map(member => member.trim());

    const toSend = {
      groupID: id,
      groupName: name,
      groupMembers: groupMembers,
    };
    const outer = this;
    const instance = axios.create({timeout: 10000});
    console.log(toSend);
    instance.defaults.headers.common['Content-Type'] = 'application/json';
    instance
      .post('http://' + window.location.hostname + '/group/create', toSend)
      .then(res => {
        console.log(res);
        outer.setState({groupResponse: res});
      })
      .catch(err => {
        console.log(err);
        outer.setState({groupResponse: err});
      });
  };
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

  render() {
    const {response, groupResponse} = this.state;
    console.log(groupResponse);
    return (
      <div>
        <MenuBar />
        <br />
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

        <Card>
          <Card.Body>
            <Card.Title>Create A Group</Card.Title>
            <Form>
              <Form.Group>
                <Col xs="5">
                  <Form.Label>Group ID</Form.Label>
                  <Form.Control type="text" id="groupID" />
                </Col>
              </Form.Group>

              <Form.Group>
                <Col xs="5">
                  <Form.Label>Group Name</Form.Label>
                  <Form.Control type="text" id="groupName" />
                </Col>
              </Form.Group>

              <Form.Group>
                <Col xs="5">
                  <Form.Label>Group Members</Form.Label>
                  <Form.Control type="text" id="groupMembers" />
                </Col>
              </Form.Group>

              <Col>
                <Button
                  variant="success"
                  type="button"
                  onClick={e => this.handleSubmitGroup(e)}
                >
                  Submit Group
                </Button>
                <div>
                  <br />
                  {groupResponse.data}
                </div>
              </Col>
            </Form>
          </Card.Body>
        </Card>
      </div>
    );
  }
}

export default HomePage;
