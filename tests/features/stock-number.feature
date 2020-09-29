Feature: ID to stock number

#  Scenario: I expect to get a valid stock number for an existing carlaed
#    Given I expect to receive following response from carlead-service with carId 12345
#      """
#      {"id":12345,"stockNumber":"AB12345"}
#      """
#    When I send "GET" request to "/car/12345/stock-number"
#    Then the response code should be 200
#    And the response should match json:
#      """
#      {
#        "12345": "AB12345"
#      }
#      """

#  Scenario: I expect to get 404 when I ask for a non-existing carlaed
#    When I send "GET" request to "/car/54321/stock-numebr"
#    Then the response code should be 404
