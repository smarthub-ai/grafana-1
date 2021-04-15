import { getTheme } from '@grafana/ui';
import { ThemeChangedEvent } from 'app/types/events';
import appEvents from '../app_events';
import { config } from '../config';
import { PreferencesService } from './PreferencesService';

export async function toggleTheme(runtimeOnly: boolean) {
  const currentTheme = config.theme;
  const newTheme = getTheme(currentTheme.isDark ? 'light' : 'dark');
  appEvents.publish(new ThemeChangedEvent(newTheme));

  if (runtimeOnly) {
    return;
  }

  // Add css file for new theme
  const newCssLink = document.createElement('link');
  newCssLink.rel = 'stylesheet';
  newCssLink.href = config.bootData.themePaths[newTheme.type];
  document.body.appendChild(newCssLink);

  // Remove old css file
  const bodyLinks = document.getElementsByTagName('link');
  for (let i = 0; i < bodyLinks.length; i++) {
    const link = bodyLinks[i];

    if (link.href && link.href.indexOf(`build/grafana.${currentTheme.type}`) > 0) {
      // Remove existing link after a 500ms to allow new css to load to avoid flickering
      // If we add new css at the same time we remove current one the page will be rendered without css
      // As the new css file is loading
      setTimeout(() => link.remove(), 500);
    }
  }

  // Persist new theme
  const service = new PreferencesService('user');
  const currentPref = await service.load();

  await service.update({
    ...currentPref,
    theme: newTheme.type,
  });
}
