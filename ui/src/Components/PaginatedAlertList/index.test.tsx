import { mount } from "enzyme";

import { useFetchGetMock } from "__fixtures__/useFetchGet";
import { AlertStore } from "Stores/AlertStore";
import { useFetchGet } from "Hooks/useFetchGet";
import { PaginatedAlertList } from ".";

let alertStore: AlertStore;

beforeEach(() => {
  jest.useFakeTimers();
  jest.setSystemTime(new Date(Date.UTC(2000, 0, 1, 0, 30, 0)));

  alertStore = new AlertStore([]);

  alertStore.data.setUpstreams({
    counters: { total: 1, healthy: 1, failed: 0 },
    instances: [
      {
        name: "am1",
        cluster: "am",
        uri: "http://localhost:9093",
        publicURI: "http://localhost:9093",
        readonly: false,
        error: "",
        version: "0.24.0",
        headers: {},
        corsCredentials: "include",
        clusterMembers: ["am1"],
      },
    ],
    clusters: { am: ["am1"] },
  });
});

afterEach(() => {
  jest.restoreAllMocks();
  jest.useRealTimers();
  document.body.className = "";
});

describe("<PaginatedAlertList />", () => {
  it("renders Placeholder while loading preview", () => {
    useFetchGetMock.fetch.setMockedData({
      response: null,
      error: null,
      isLoading: true,
      isRetrying: false,
      retryCount: 0,
      get: jest.fn(),
      cancelGet: jest.fn(),
    });
    const tree = mount(
      <PaginatedAlertList alertStore={alertStore} filters={["foo=bar"]} />
    );
    expect(tree.find("Placeholder")).toHaveLength(1);
  });

  it("renders Placeholder while response is empty", () => {
    useFetchGetMock.fetch.setMockedData({
      response: null,
      error: null,
      isLoading: false,
      isRetrying: false,
      retryCount: 0,
      get: jest.fn(),
      cancelGet: jest.fn(),
    });
    const tree = mount(
      <PaginatedAlertList alertStore={alertStore} filters={["foo=bar"]} />
    );
    expect(tree.find("Placeholder")).toHaveLength(1);
  });

  it("renders LabelSetList with StaticLabel on mount", () => {
    useFetchGetMock.fetch.setMockedData({
      response: {
        alerts: [
          [
            { name: "alertname", value: "Fake Alert" },
            { name: "foo", value: "1" },
            { name: "bar", value: "2" },
          ],
        ],
      },
      error: undefined,
      isLoading: false,
      isRetrying: false,
      retryCount: 0,
      get: jest.fn(),
      cancelGet: jest.fn(),
    });
    const tree = mount(
      <PaginatedAlertList alertStore={alertStore} filters={["foo=bar"]} />
    );
    expect(tree.find("LabelSetList")).toHaveLength(1);
    expect(tree.find("Memo(StaticLabel)")).toHaveLength(3);
  });

  it("renders empty LabelSetList with empty response", () => {
    useFetchGetMock.fetch.setMockedData({
      response: { alerts: [] },
      error: null,
      isLoading: false,
      isRetrying: false,
      retryCount: 0,
      get: jest.fn(),
      cancelGet: jest.fn(),
    });
    const tree = mount(
      <PaginatedAlertList alertStore={alertStore} filters={["foo=bar"]} />
    );
    expect(tree.find("LabelSetList")).toHaveLength(1);
    expect(tree.find("Memo(StaticLabel)")).toHaveLength(0);
  });

  it("fetches affected alerts on mount", () => {
    mount(<PaginatedAlertList alertStore={alertStore} filters={["foo=bar"]} />);
    expect(useFetchGet).toHaveBeenCalled();
  });

  it("renders StaticLabel after fetch", () => {
    useFetchGetMock.fetch.setMockedData({
      response: {
        alerts: [
          [
            { name: "alertname", value: "Fake Alert" },
            { name: "foo", value: "1" },
            { name: "bar", value: "2" },
          ],
        ],
      },
      error: undefined,
      isLoading: false,
      isRetrying: false,
      retryCount: 0,
      get: jest.fn(),
      cancelGet: jest.fn(),
    });
    const tree = mount(
      <PaginatedAlertList
        alertStore={alertStore}
        filters={["foo=bar"]}
        title="Affected alerts"
      />
    );
    expect(tree.text()).toMatch(/Affected alerts/);
    expect(tree.find("Memo(StaticLabel)")).toHaveLength(3);
  });

  it("handles empty grid response correctly", () => {
    useFetchGetMock.fetch.setMockedData({
      response: { alerts: [] },
      error: null,
      isLoading: false,
      isRetrying: false,
      retryCount: 0,
      get: jest.fn(),
      cancelGet: jest.fn(),
    });
    const tree = mount(
      <PaginatedAlertList alertStore={alertStore} filters={["foo=bar"]} />
    );
    expect(tree.text()).toMatch(/No alerts matched/);
  });

  it("renders FetchError on failed preview fetch", () => {
    useFetchGetMock.fetch.setMockedData({
      response: null,
      error: "fake error",
      isLoading: false,
      isRetrying: false,
      retryCount: 0,
      get: jest.fn(),
      cancelGet: jest.fn(),
    });
    const tree = mount(
      <PaginatedAlertList alertStore={alertStore} filters={["foo=bar"]} />
    );
    expect(tree.find("FetchError")).toHaveLength(1);
  });
});
