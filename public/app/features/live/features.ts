import { LiveChannelConfig } from '@grafana/data';
import { getDashboardChannelsFeature } from './dashboard/dashboardWatcher';
import { LiveMeasurementsSupport } from './measurements/measurementsSupport';
import { grafanaLiveCoreFeatures } from './scopes';

export function registerLiveFeatures() {
  const channels: LiveChannelConfig[] = [
    {
      path: 'random-2s-stream',
      description: 'Random stream with points every 2s',
    },
    {
      path: 'random-flakey-stream',
      description: 'Random stream with flakey data points',
    },
    {
      path: 'random-20Hz-stream',
      description: 'Random stream with points in 20Hz',
    },
  ];

  grafanaLiveCoreFeatures.register({
    name: 'testdata',
    support: {
      getChannelConfig: (path: string) => {
        return channels.find((c) => c.path === path);
      },
      getSupportedPaths: () => channels,
    },
    description: 'Test data generations',
  });

  const broadcastConfig: LiveChannelConfig = {
    path: '${path}',
    description: 'Broadcast any messages to a channel',
    canPublish: () => true,
  };

  grafanaLiveCoreFeatures.register({
    name: 'broadcast',
    support: {
      getChannelConfig: (path: string) => {
        return broadcastConfig;
      },
      getSupportedPaths: () => [broadcastConfig],
    },
    description: 'Broadcast will send/receive any JSON object in a channel',
  });

  grafanaLiveCoreFeatures.register({
    name: 'measurements',
    support: new LiveMeasurementsSupport(),
    description: 'These channels listen for measurements and produce DataFrames',
  });

  // dashboard/*
  grafanaLiveCoreFeatures.register(getDashboardChannelsFeature());
}
