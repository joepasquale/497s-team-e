import React from 'react';
import {Button} from 'react-bootstrap';
import {Form, Col, Card} from 'react-bootstrap';
import axios from 'axios';

export class HomePage extends React.Component {
  constructor(props) {
    super(props);
    this.mocking = true;
    this.data = this.mocking
      ? {
          eventID: 0,
          groupID: 0,
          eventName: '',
          eventTime: '',
          eventLocation: '',
        }
      : {};
    this.state = {
      eventCreateResponse: '',
      eventReadResponse: '',
      eventUpdateResponse: '',
      eventDeleteResponse: '',
    };
  }
  handleSubmitEventCreate = () => {
    const id = document.getElementById('eventCreateID').value;
    const groupID = document.getElementById('eventCreateGroupID').value;
    const name = document.getElementById('eventCreateName').value;
    const time = document.getElementById('eventCreateTime').value;
    const location = document.getElementById('eventCreateLocation').value;
    const toSend = {
      eventID: id,
      groupID: groupID,
      eventName: name,
      eventTime: time,
      eventLocation: location,
    };
    const outer = this;
    const instance = axios.create({timeout: 10000});
    console.log(toSend);
    instance.defaults.headers.common['Content-Type'] = 'application/json';
    instance
      .post('http://' + window.location.hostname + '/event/create', toSend)
      .then(res => {
        console.log(res);
        outer.setState({eventCreateResponse: res});
      })
      .catch(err => {
        console.log(err);
        outer.setState({eventCreateResponse: err});
      });
  };
  handleSubmitEventRead = () => {
    const id = document.getElementById('eventReadID').value;

    const toSend = {
      eventID: id,
    };
    const outer = this;
    const instance = axios.create({timeout: 10000});
    console.log(toSend);
    instance.defaults.headers.common['Content-Type'] = 'application/json';
    instance
      .post('http://' + window.location.hostname + '/event/read', toSend)
      .then(res => {
        console.log(res);
        outer.setState({eventReadResponse: res});
      })
      .catch(err => {
        console.log(err);
        outer.setState({eventReadResponse: err});
      });
  };
  handleSubmitEventUpdate = () => {
    const id = document.getElementById('eventUpdateID').value;
    const name = document.getElementById('eventUpdateName').value;
    const time = document.getElementById('eventUpdateTime').value;
    const location = document.getElementById('eventUpdateLocation').value;
    const toSend = {
      eventID: id,
      eventName: name,
      eventTime: time,
      eventLocation: location,
    };
    const outer = this;
    const instance = axios.create({timeout: 10000});
    console.log(toSend);
    instance.defaults.headers.common['Content-Type'] = 'application/json';
    instance
      .post('http://' + window.location.hostname + '/event/update', toSend)
      .then(res => {
        console.log(res);
        outer.setState({eventUpdateResponse: res});
      })
      .catch(err => {
        console.log(err);
        outer.setState({eventUpdateResponse: err});
      });
  };

  handleSubmitEventDelete = () => {
    const id = document.getElementById('eventDeleteID').value;

    const toSend = {
      eventID: id,
    };
    const outer = this;
    const instance = axios.create({timeout: 10000});
    console.log(toSend);
    instance.defaults.headers.common['Content-Type'] = 'application/json';
    instance
      .post('http://' + window.location.hostname + '/event/delete', toSend)
      .then(res => {
        console.log(res);
        outer.setState({eventDeleteResponse: res});
      })
      .catch(err => {
        console.log(err);
        outer.setState({eventDeleteResponse: err});
      });
  };

  render() {
    const {
      eventCreateResponse,
      eventReadResponse,
      eventUpdateResponse,
      eventDeleteResponse,
    } = this.state;

    const createEvent = (
      <Card>
        <Card.Body>
          <Card.Title>Create an Event</Card.Title>
          <Form>
            <Form.Group>
              <Col xs="5">
                <Form.Label>Event ID</Form.Label>
                <Form.Control type="text" id="eventCreateID" />
              </Col>
            </Form.Group>

            <Form.Group>
              <Col xs="5">
                <Form.Label>Group ID</Form.Label>
                <Form.Control type="text" id="eventCreateGroupID" />
              </Col>
            </Form.Group>

            <Form.Group>
              <Col xs="5">
                <Form.Label>Event Name</Form.Label>
                <Form.Control type="text" id="eventCreateName" />
              </Col>
            </Form.Group>

            <Form.Group>
              <Col xs="5">
                <Form.Label>Event Time</Form.Label>
                <Form.Control type="text" id="eventCreateTime" />
              </Col>
            </Form.Group>

            <Form.Group>
              <Col xs="5">
                <Form.Label>Event Location</Form.Label>
                <Form.Control type="text" id="eventCreateLocation" />
              </Col>
            </Form.Group>

            <Col>
              <Button
                variant="success"
                type="button"
                onClick={e => this.handleSubmitEventCreate(e)}
              >
                Create Event
              </Button>
              <div>
                <br />
                {eventCreateResponse.data}
              </div>
            </Col>
          </Form>
        </Card.Body>
      </Card>
    );

    const readEvent = (
      <Card>
        <Card.Body>
          <Card.Title>Read an Event</Card.Title>
          <Form>
            <Form.Group>
              <Col xs="5">
                <Form.Label>Event ID</Form.Label>
                <Form.Control type="text" id="eventReadID" />
              </Col>
            </Form.Group>

            <Col>
              <Button
                variant="success"
                type="button"
                onClick={e => this.handleSubmitEventRead(e)}
              >
                Read Event
              </Button>
              <div>
                <br />
                {eventReadResponse.data}
              </div>
            </Col>
          </Form>
        </Card.Body>
      </Card>
    );

    const updateEvent = (
      <Card>
        <Card.Body>
          <Card.Title>Update an Event</Card.Title>
          <Form>
            <Form.Group>
              <Col xs="5">
                <Form.Label>Event ID</Form.Label>
                <Form.Control type="text" id="eventUpdateID" />
              </Col>
            </Form.Group>

            <Form.Group>
              <Col xs="5">
                <Form.Label>Event Name</Form.Label>
                <Form.Control type="text" id="eventUpdateName" />
              </Col>
            </Form.Group>

            <Form.Group>
              <Col xs="5">
                <Form.Label>Event Time</Form.Label>
                <Form.Control type="text" id="eventUpdateTime" />
              </Col>
            </Form.Group>

            <Form.Group>
              <Col xs="5">
                <Form.Label>Event Location</Form.Label>
                <Form.Control type="text" id="eventUpdateLocation" />
              </Col>
            </Form.Group>

            <Col>
              <Button
                variant="success"
                type="button"
                onClick={e => this.handleSubmitEventUpdate(e)}
              >
                Update Event
              </Button>
              <div>
                <br />
                {eventUpdateResponse.data}
              </div>
            </Col>
          </Form>
        </Card.Body>
      </Card>
    );

    const deleteEvent = (
      <Card>
        <Card.Body>
          <Card.Title>Delete an Event</Card.Title>
          <Form>
            <Form.Group>
              <Col xs="5">
                <Form.Label>Event ID</Form.Label>
                <Form.Control type="text" id="eventDeleteID" />
              </Col>
            </Form.Group>

            <Col>
              <Button
                variant="success"
                type="button"
                onClick={e => this.handleSubmitEventDelete(e)}
              >
                Delete Event
              </Button>
              <div>
                <br />
                {eventDeleteResponse.data}
              </div>
            </Col>
          </Form>
        </Card.Body>
      </Card>
    );

    return (
      <div>
        {createEvent}
        {readEvent}
        {updateEvent}
        {deleteEvent}
      </div>
    );
  }
}

export default HomePage;
