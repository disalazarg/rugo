describe Worker do
  it "produces a JSON string" do
    json_string = Worker.login("user")
    json = JSON.parse json_string

    expect(json).to be_a Hash
  end
end
