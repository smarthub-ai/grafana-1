import {
  PanelMenuItem,
  PluginExtension,
  PluginExtensionLink,
  PluginExtensionTypes,
  PluginExtensionPlacements,
} from '@grafana/data';
import {
  PluginExtensionPanelContext,
  PluginExtensionRegistryItem,
  setPluginsExtensionRegistry,
} from '@grafana/runtime';
import config from 'app/core/config';
import * as actions from 'app/features/explore/state/main';
import { setStore } from 'app/store/store';

import { PanelModel } from '../state';
import { createDashboardModelFixture } from '../state/__fixtures__/dashboardFixtures';

import { getPanelMenu } from './getPanelMenu';

jest.mock('app/core/services/context_srv', () => ({
  contextSrv: {
    hasAccessToExplore: () => true,
  },
}));

describe('getPanelMenu()', () => {
  beforeEach(() => {
    setPluginsExtensionRegistry({});
  });

  it('should return the correct panel menu items', () => {
    const panel = new PanelModel({});
    const dashboard = createDashboardModelFixture({});

    const menuItems = getPanelMenu(dashboard, panel);
    expect(menuItems).toMatchInlineSnapshot(`
      [
        {
          "iconClassName": "eye",
          "onClick": [Function],
          "shortcut": "v",
          "text": "View",
        },
        {
          "iconClassName": "edit",
          "onClick": [Function],
          "shortcut": "e",
          "text": "Edit",
        },
        {
          "iconClassName": "share-alt",
          "onClick": [Function],
          "shortcut": "p s",
          "text": "Share",
        },
        {
          "iconClassName": "compass",
          "onClick": [Function],
          "shortcut": "x",
          "text": "Explore",
        },
        {
          "iconClassName": "info-circle",
          "onClick": [Function],
          "shortcut": "i",
          "subMenu": [
            {
              "onClick": [Function],
              "text": "Panel JSON",
            },
          ],
          "text": "Inspect",
          "type": "submenu",
        },
        {
          "iconClassName": "cube",
          "onClick": [Function],
          "subMenu": [
            {
              "onClick": [Function],
              "shortcut": "p d",
              "text": "Duplicate",
            },
            {
              "onClick": [Function],
              "text": "Copy",
            },
            {
              "onClick": [Function],
              "text": "Create library panel",
            },
          ],
          "text": "More...",
          "type": "submenu",
        },
        {
          "text": "",
          "type": "divider",
        },
        {
          "iconClassName": "trash-alt",
          "onClick": [Function],
          "shortcut": "p r",
          "text": "Remove",
        },
      ]
    `);
  });

  describe('when extending panel menu from plugins', () => {
    it('should contain menu item from link extension', () => {
      setPluginsExtensionRegistry({
        [PluginExtensionPlacements.DashboardPanelMenu]: [
          createRegistryItem<PluginExtensionLink>({
            type: PluginExtensionTypes.link,
            title: 'Declare incident',
            description: 'Declaring an incident in the app',
            path: '/a/grafana-basic-app/declare-incident',
            key: 1,
          }),
        ],
      });

      const panel = new PanelModel({});
      const dashboard = createDashboardModelFixture({});
      const menuItems = getPanelMenu(dashboard, panel);
      const moreSubMenu = menuItems.find((i) => i.text === 'Extensions')?.subMenu;

      expect(moreSubMenu).toEqual(
        expect.arrayContaining([
          expect.objectContaining({
            text: 'Declare incident',
            href: '/a/grafana-basic-app/declare-incident',
          }),
        ])
      );
    });

    it('should truncate menu item title to 25 chars', () => {
      setPluginsExtensionRegistry({
        [PluginExtensionPlacements.DashboardPanelMenu]: [
          createRegistryItem<PluginExtensionLink>({
            type: PluginExtensionTypes.link,
            title: 'Declare incident when pressing this amazing menu item',
            description: 'Declaring an incident in the app',
            path: '/a/grafana-basic-app/declare-incident',
            key: 1,
          }),
        ],
      });

      const panel = new PanelModel({});
      const dashboard = createDashboardModelFixture({});
      const menuItems = getPanelMenu(dashboard, panel);
      const moreSubMenu = menuItems.find((i) => i.text === 'Extensions')?.subMenu;

      expect(moreSubMenu).toEqual(
        expect.arrayContaining([
          expect.objectContaining({
            text: 'Declare incident when...',
            href: '/a/grafana-basic-app/declare-incident',
          }),
        ])
      );
    });

    it('should use extension for panel menu returned by configure function', () => {
      const configure: PluginExtensionRegistryItem<PluginExtensionLink> = () => ({
        title: 'Wohoo',
        type: PluginExtensionTypes.link,
        description: 'Declaring an incident in the app',
        path: '/a/grafana-basic-app/declare-incident',
        key: 1,
      });

      setPluginsExtensionRegistry({
        [PluginExtensionPlacements.DashboardPanelMenu]: [
          createRegistryItem<PluginExtensionLink>(
            {
              type: PluginExtensionTypes.link,
              title: 'Declare incident when pressing this amazing menu item',
              description: 'Declaring an incident in the app',
              path: '/a/grafana-basic-app/declare-incident',
              key: 1,
            },
            configure
          ),
        ],
      });

      const panel = new PanelModel({});
      const dashboard = createDashboardModelFixture({});
      const menuItems = getPanelMenu(dashboard, panel);
      const moreSubMenu = menuItems.find((i) => i.text === 'Extensions')?.subMenu;

      expect(moreSubMenu).toEqual(
        expect.arrayContaining([
          expect.objectContaining({
            text: 'Wohoo',
            href: '/a/grafana-basic-app/declare-incident',
          }),
        ])
      );
    });

    it('should hide menu item if configure function returns undefined', () => {
      setPluginsExtensionRegistry({
        [PluginExtensionPlacements.DashboardPanelMenu]: [
          createRegistryItem<PluginExtensionLink>(
            {
              type: PluginExtensionTypes.link,
              title: 'Declare incident when pressing this amazing menu item',
              description: 'Declaring an incident in the app',
              path: '/a/grafana-basic-app/declare-incident',
              key: 1,
            },
            () => undefined
          ),
        ],
      });

      const panel = new PanelModel({});
      const dashboard = createDashboardModelFixture({});
      const menuItems = getPanelMenu(dashboard, panel);
      const moreSubMenu = menuItems.find((i) => i.text === 'Extensions')?.subMenu;

      expect(moreSubMenu).toEqual(
        expect.not.arrayContaining([
          expect.objectContaining({
            text: 'Declare incident when...',
            href: '/a/grafana-basic-app/declare-incident',
          }),
        ])
      );
    });

    it('should pass context with correct values when configuring extension', () => {
      const configure = jest.fn();

      setPluginsExtensionRegistry({
        [PluginExtensionPlacements.DashboardPanelMenu]: [
          createRegistryItem<PluginExtensionLink>(
            {
              type: PluginExtensionTypes.link,
              title: 'Declare incident when pressing this amazing menu item',
              description: 'Declaring an incident in the app',
              path: '/a/grafana-basic-app/declare-incident',
              key: 1,
            },
            configure
          ),
        ],
      });

      const panel = new PanelModel({
        type: 'timeseries',
        id: 1,
        title: 'My panel',
        targets: [
          {
            refId: 'A',
            datasource: {
              type: 'testdata',
            },
          },
        ],
      });

      const dashboard = createDashboardModelFixture({
        timezone: 'utc',
        time: {
          from: 'now-5m',
          to: 'now',
        },
        tags: ['database', 'panel'],
        uid: '123',
        title: 'My dashboard',
      });

      getPanelMenu(dashboard, panel);

      const context: PluginExtensionPanelContext = {
        pluginId: 'timeseries',
        id: 1,
        title: 'My panel',
        timeZone: 'utc',
        timeRange: {
          from: 'now-5m',
          to: 'now',
        },
        targets: [
          {
            refId: 'A',
            pluginId: 'testdata',
          },
        ],
        dashboard: {
          tags: ['database', 'panel'],
          uid: '123',
          title: 'My dashboard',
        },
      };

      expect(configure).toBeCalledWith(context);
    });

    it('should pass context that can not be edited in configure function', () => {
      const configure: PluginExtensionRegistryItem<PluginExtensionLink> = (context) => {
        // trying to change values in the context
        // @ts-ignore
        context.pluginId = 'changed';

        return {
          type: PluginExtensionTypes.link,
          title: 'Declare incident when pressing this amazing menu item',
          description: 'Declaring an incident in the app',
          path: '/a/grafana-basic-app/declare-incident',
          key: 1,
        };
      };

      setPluginsExtensionRegistry({
        [PluginExtensionPlacements.DashboardPanelMenu]: [
          createRegistryItem<PluginExtensionLink>(
            {
              type: PluginExtensionTypes.link,
              title: 'Declare incident when pressing this amazing menu item',
              description: 'Declaring an incident in the app',
              path: '/a/grafana-basic-app/declare-incident',
              key: 1,
            },
            configure
          ),
        ],
      });

      const panel = new PanelModel({
        type: 'timeseries',
        id: 1,
        title: 'My panel',
        targets: [
          {
            refId: 'A',
            datasource: {
              type: 'testdata',
            },
          },
        ],
      });

      const dashboard = createDashboardModelFixture({
        timezone: 'utc',
        time: {
          from: 'now-5m',
          to: 'now',
        },
        tags: ['database', 'panel'],
        uid: '123',
        title: 'My dashboard',
      });

      expect(() => getPanelMenu(dashboard, panel)).toThrowError(TypeError);
    });
  });

  describe('when panel is in view mode', () => {
    it('should return the correct panel menu items', () => {
      const getExtendedMenu = () => [{ text: 'Toggle legend', shortcut: 'p l', click: jest.fn() }];
      const ctrl: any = { getExtendedMenu };
      const scope: any = { $$childHead: { ctrl } };
      const angularComponent: any = { getScope: () => scope };
      const panel = new PanelModel({ isViewing: true });
      const dashboard = createDashboardModelFixture({});

      const menuItems = getPanelMenu(dashboard, panel, angularComponent);
      expect(menuItems).toMatchInlineSnapshot(`
        [
          {
            "iconClassName": "eye",
            "onClick": [Function],
            "shortcut": "v",
            "text": "View",
          },
          {
            "iconClassName": "edit",
            "onClick": [Function],
            "shortcut": "e",
            "text": "Edit",
          },
          {
            "iconClassName": "share-alt",
            "onClick": [Function],
            "shortcut": "p s",
            "text": "Share",
          },
          {
            "iconClassName": "compass",
            "onClick": [Function],
            "shortcut": "x",
            "text": "Explore",
          },
          {
            "iconClassName": "info-circle",
            "onClick": [Function],
            "shortcut": "i",
            "subMenu": [
              {
                "onClick": [Function],
                "text": "Panel JSON",
              },
            ],
            "text": "Inspect",
            "type": "submenu",
          },
          {
            "iconClassName": "cube",
            "onClick": [Function],
            "subMenu": [
              {
                "href": undefined,
                "onClick": [Function],
                "shortcut": "p l",
                "text": "Toggle legend",
              },
            ],
            "text": "More...",
            "type": "submenu",
          },
        ]
      `);
    });
  });

  describe('onNavigateToExplore', () => {
    const testSubUrl = '/testSubUrl';
    const testUrl = '/testUrl';
    const windowOpen = jest.fn();
    let event: any;
    let explore: PanelMenuItem;
    let navigateSpy: any;

    beforeAll(() => {
      const panel = new PanelModel({});
      const dashboard = createDashboardModelFixture({});
      const menuItems = getPanelMenu(dashboard, panel);
      explore = menuItems.find((item) => item.text === 'Explore') as PanelMenuItem;
      navigateSpy = jest.spyOn(actions, 'navigateToExplore');
      window.open = windowOpen;

      event = {
        ctrlKey: true,
        preventDefault: jest.fn(),
      };

      setStore({ dispatch: jest.fn() } as any);
    });

    it('should navigate to url without subUrl', () => {
      explore.onClick!(event);

      const openInNewWindow = navigateSpy.mock.calls[0][1].openInNewWindow;

      openInNewWindow(testUrl);

      expect(windowOpen).toHaveBeenLastCalledWith(testUrl);
    });

    it('should navigate to url with subUrl', () => {
      config.appSubUrl = testSubUrl;
      explore.onClick!(event);

      const openInNewWindow = navigateSpy.mock.calls[0][1].openInNewWindow;

      openInNewWindow(testUrl);

      expect(windowOpen).toHaveBeenLastCalledWith(`${testSubUrl}${testUrl}`);
    });
  });
});

function createRegistryItem<T extends PluginExtension, C extends object = object>(
  extension: T,
  configure?: PluginExtensionRegistryItem<T, C>
): PluginExtensionRegistryItem<T, C> {
  const defaultConfigure = () => extension;
  return configure || defaultConfigure;
}
