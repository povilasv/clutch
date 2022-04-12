import * as React from "react";
import renderer from "react-test-renderer";

import useDataLayoutManager from "../manager";

describe("Manager Default State", () => {
  it("does not share an initial state reference", () => {
    let manager;

    function TestComponent() {
      manager = useDataLayoutManager({
        a: {},
        b: {},
      });

      return null;
    }

    renderer.create(<TestComponent />);

    manager.state.a.data = { foo: "bar" };
    expect(manager.state.b.data).toEqual({});
  });
});
