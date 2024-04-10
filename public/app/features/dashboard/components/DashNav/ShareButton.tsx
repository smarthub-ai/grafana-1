import React from 'react';

import { selectors as e2eSelectors } from '@grafana/e2e-selectors/src';
import { locationService } from '@grafana/runtime';
import { Button } from '@grafana/ui';
import { Trans } from 'app/core/internationalization';
import { DashboardModel } from 'app/features/dashboard/state';
import { DashboardInteractions } from 'app/features/dashboard-scene/utils/interactions';

export const ShareButton = ({ dashboard }: { dashboard: DashboardModel }) => {
  return (
    <Button
      data-testid={e2eSelectors.pages.Dashboard.DashNav.shareButton}
      variant="primary"
      size="sm"
      onClick={() => {
        DashboardInteractions.toolbarShareClick();
        locationService.partial({ shareView: 'link' });
      }}
    >
      <Trans i18nKey="dashboard.toolbar.share-button">Share</Trans>
    </Button>
  );
};
