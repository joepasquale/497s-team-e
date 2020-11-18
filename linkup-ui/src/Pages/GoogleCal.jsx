import React from 'react';
import ReactDOM from 'react-dom'
import { Button } from 'react-bootstrap';
import { Form, Col, Card } from 'react-bootstrap';
import axios from 'axios';
import GoogleLogin from 'react-google-login'

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
  handleGoogleLogin = (response) => {
    const toSend = {
      "provider": "google-oauth2",
      "code": response.code,
      "redirect-uri": "postmessage"
    };
    const outer = this;
    const instance = axios.create({ timeout: 10000 });
    console.log(toSend);
    instance.defaults.headers.common['Content-Type'] = 'application/json';
    instance
      .post('http://' + window.location.hostname + '/gcal/auth', toSend)
      .then(res => {
        console.log(res);
        outer.setState({ eventExportResponse: res });
      })
      .catch(err => {
        console.log(err);
        outer.setState({ eventExportResponse: err });
      });
  };

  handleSubmitEventExport = () => {
    const id = document.getElementById('eventExportID').value;

    const toSend = {
      eventID: id,
    };
    const outer = this;
    const instance = axios.create({ timeout: 10000 });
    console.log(toSend);
    instance.defaults.headers.common['Content-Type'] = 'application/json';
    instance
      .post('http://' + window.location.hostname + '/gcal/add', toSend)
      .then(res => {
        console.log(res);
        outer.setState({ eventExportResponse: res });
      })
      .catch(err => {
        console.log(err);
        outer.setState({ eventExportResponse: err });
      });
  };

  render() {
    const {
      eventExportResponse,
    } = this.state;

    const responseGoogle = (response) => {
      this.handleGoogleLogin(response);
    }

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
              <GoogleLogin
                clientId="350429252210-hq617ss9idkeat0h66hbop59ul53mpnf.apps.googleusercontent.com"
                buttonText="Login"
                responseType="code"
                //on success post to backend
                onSuccess={responseGoogle}
                onFailure={responseGoogle}
                cookiePolicy={'single_host_origin'} />
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
