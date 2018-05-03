Feature: connect with worker

  Scenario: Receive connect request from worker
    Given a worker
    When worker try to connect sending "alohomora" message
    Then server should respond "imperio" message
