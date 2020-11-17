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
      eventExportResponse: '',
    };
  }
  handleSubmitEventExport = () => {
    const id = document.getElementById('eventExportID').value;

    const toSend = {
      eventID: id,
    };
    const outer = this;
    const instance = axios.create({timeout: 10000});
    console.log(toSend);
    instance.defaults.headers.common['Content-Type'] = 'application/json';
    instance
      .post('http://' + window.location.hostname + '/gcal/add', toSend)
      .then(res => {
        console.log(res);
        outer.setState({eventExportResponse: res});
      })
      .catch(err => {
        console.log(err);
        outer.setState({eventExportResponse: err});
      });
  };

  render() {
    const {
      eventExportResponse,
    } = this.state;

    const exportEvent = (
      <Card>
        <Card.Body>
          <Card.Title>Export to GCal</Card.Title>
          <Form>
            <Form.Group>
              <Col xs="5">
                <Form.Label>Event ID</Form.Label>
                <Form.Control type="text" id="eventExportID" />
              </Col>
            </Form.Group>

            <Col>
              <Button
                variant="success"
                type="button"
                onClick={e => this.handleSubmitEventExport(e)}
              >
                Export to Google Calendar
              </Button>
              <div>
                <br />
                {eventExportResponse.data}
              </div>
            </Col>
          </Form>
        </Card.Body>
      </Card>
    );

    return (
      <div>
        {exportEvent}
      </div>
    );
  }
}

export default HomePage;
