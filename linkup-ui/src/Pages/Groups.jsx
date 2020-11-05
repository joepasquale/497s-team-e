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
          groupID: '',
          groupName: '',
          groupMembers: '',
        }
      : {};
    this.state = {
      groupCreateResponse: '',
      groupReadResponse: '',
      groupUpdateResponse: '',
      groupDeleteResponse: '',
    };
  }
  handleSubmitGroupCreate = () => {
    const id = document.getElementById('groupCreateID').value;
    const name = document.getElementById('groupCreateName').value;
    const members = document.getElementById('groupCreateMembers').value;

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
        outer.setState({groupCreateResponse: res});
      })
      .catch(err => {
        console.log(err);
        outer.setState({groupCreateResponse: err});
      });
  };
  handleSubmitGroupRead = () => {
    const id = document.getElementById('groupReadID').value;

    const toSend = {
      groupID: id,
    };
    const outer = this;
    const instance = axios.create({timeout: 10000});
    console.log(toSend);
    instance.defaults.headers.common['Content-Type'] = 'application/json';
    instance
      .post('http://' + window.location.hostname + '/group/read', toSend)
      .then(res => {
        console.log(res);
        outer.setState({groupReadResponse: res});
      })
      .catch(err => {
        console.log(err);
        outer.setState({groupReadResponse: err});
      });
  };
  handleSubmitGroupUpdate = () => {
    const id = document.getElementById('groupUpdateID').value;
    const name = document.getElementById('groupUpdateName').value;
    const members = document.getElementById('groupUpdateMembers').value;

    const groupMembers = members.split(',').map(member => member.trim());

    const toSend = {
      groupID: id,
      groupName: name,
      groupMembers: groupMembers,
    };
    const outer = this;
    const instance = axios.create({timeout: 10000});
    instance.defaults.headers.common['Content-Type'] = 'application/json';
    instance
      .post('http://' + window.location.hostname + '/group/update', toSend)
      .then(res => {
        console.log(res);
        outer.setState({groupUpdateResponse: res});
      })
      .catch(err => {
        console.log(err);
        outer.setState({groupUpdateResponse: err});
      });
  };

  handleSubmitGroupDelete = () => {
    const id = document.getElementById('groupDeleteID').value;

    const toSend = {
      groupID: id,
    };
    const outer = this;
    const instance = axios.create({timeout: 10000});
    console.log(toSend);
    instance.defaults.headers.common['Content-Type'] = 'application/json';
    instance
      .post('http://' + window.location.hostname + '/group/delete', toSend)
      .then(res => {
        console.log(res);
        outer.setState({groupDeleteResponse: res});
      })
      .catch(err => {
        console.log(err);
        outer.setState({groupDeleteResponse: err});
      });
  };

  render() {
    const {
      groupCreateResponse,
      groupReadResponse,
      groupUpdateResponse,
      groupDeleteResponse,
    } = this.state;

    const createGroup = (
      <Card>
        <Card.Body>
          <Card.Title>Create a Group</Card.Title>
          <Form>
            <Form.Group>
              <Col xs="5">
                <Form.Label>Group ID</Form.Label>
                <Form.Control type="text" id="groupCreateID" />
              </Col>
            </Form.Group>

            <Form.Group>
              <Col xs="5">
                <Form.Label>Group Name</Form.Label>
                <Form.Control type="text" id="groupCreateName" />
              </Col>
            </Form.Group>

            <Form.Group>
              <Col xs="5">
                <Form.Label>Group Members</Form.Label>
                <Form.Control type="text" id="groupCreateMembers" />
              </Col>
            </Form.Group>

            <Col>
              <Button
                variant="success"
                type="button"
                onClick={e => this.handleSubmitGroupCreate(e)}
              >
                Create Group
              </Button>
              <div>
                <br />
                {groupCreateResponse.data}
              </div>
            </Col>
          </Form>
        </Card.Body>
      </Card>
    );

    const readGroup = (
      <Card>
        <Card.Body>
          <Card.Title>Read a Group</Card.Title>
          <Form>
            <Form.Group>
              <Col xs="5">
                <Form.Label>Group ID</Form.Label>
                <Form.Control type="text" id="groupReadID" />
              </Col>
            </Form.Group>

            <Col>
              <Button
                variant="success"
                type="button"
                onClick={e => this.handleSubmitGroupRead(e)}
              >
                Read Group
              </Button>
              <div>
                <br />
                {groupReadResponse.data}
              </div>
            </Col>
          </Form>
        </Card.Body>
      </Card>
    );

    const updateGroup = (
      <Card>
        <Card.Body>
          <Card.Title>Update a Group</Card.Title>
          <Form>
            <Form.Group>
              <Col xs="5">
                <Form.Label>Group ID</Form.Label>
                <Form.Control type="text" id="groupUpdateID" />
              </Col>
            </Form.Group>

            <Form.Group>
              <Col xs="5">
                <Form.Label>Group Name</Form.Label>
                <Form.Control type="text" id="groupUpdateName" />
              </Col>
            </Form.Group>

            <Form.Group>
              <Col xs="5">
                <Form.Label>Group Members</Form.Label>
                <Form.Control type="text" id="groupUpdateMembers" />
              </Col>
            </Form.Group>

            <Col>
              <Button
                variant="success"
                type="button"
                onClick={e => this.handleSubmitGroupUpdate(e)}
              >
                Update Group
              </Button>
              <div>
                <br />
                {groupUpdateResponse.data}
              </div>
            </Col>
          </Form>
        </Card.Body>
      </Card>
    );

    const deleteGroup = (
      <Card>
        <Card.Body>
          <Card.Title>Delete a Group</Card.Title>
          <Form>
            <Form.Group>
              <Col xs="5">
                <Form.Label>Group ID</Form.Label>
                <Form.Control type="text" id="groupDeleteID" />
              </Col>
            </Form.Group>

            <Col>
              <Button
                variant="success"
                type="button"
                onClick={e => this.handleSubmitGroupDelete(e)}
              >
                Delete Group
              </Button>
              <div>
                <br />
                {groupDeleteResponse.data}
              </div>
            </Col>
          </Form>
        </Card.Body>
      </Card>
    );

    return (
      <div>
        {createGroup}
        {readGroup}
        {updateGroup}
        {deleteGroup}
      </div>
    );
  }
}

export default HomePage;
