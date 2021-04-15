import React from 'react';
import { css, cx } from '@emotion/css';
import { GrafanaTheme } from '@grafana/data';
import { Icon } from '../Icon/Icon';
import { IconButton } from '../IconButton/IconButton';
import { HorizontalGroup } from '../Layout/Layout';
import { AlertVariant } from '../Alert/Alert';
import panelArtDark from './panelArt_dark.svg';
import panelArtLight from './panelArt_light.svg';
import { stylesFactory, useTheme } from '../../themes';

export interface InfoBoxProps extends Omit<React.HTMLAttributes<HTMLDivElement>, 'title'> {
  children: React.ReactNode;
  /** Title of the box */
  title?: string | JSX.Element;
  /** Url of the read more link */
  url?: string;
  /** Text of the read more link */
  urlTitle?: string;
  /** Indicates whether or not box should be rendered with Grafana branding background */
  branded?: boolean;
  /** Color variant of the box */
  severity?: AlertVariant;
  /** Call back to be performed when box is dismissed */
  onDismiss?: () => void;
}

/**
  @public
 */
export const InfoBox = React.memo(
  React.forwardRef<HTMLDivElement, InfoBoxProps>(
    ({ title, className, children, branded, url, urlTitle, onDismiss, severity = 'info', ...otherProps }, ref) => {
      const theme = useTheme();
      const styles = getInfoBoxStyles(theme, severity);
      const wrapperClassName = cx(branded ? styles.wrapperBranded : styles.wrapper, className);

      return (
        <div className={wrapperClassName} {...otherProps} ref={ref}>
          <div>
            <HorizontalGroup justify={'space-between'} align={'flex-start'}>
              <div>{typeof title === 'string' ? <h4>{title}</h4> : title}</div>
              {onDismiss && <IconButton name={'times'} onClick={onDismiss} />}
            </HorizontalGroup>
          </div>
          <div>{children}</div>
          {url && (
            <a href={url} className={styles.docsLink} target="_blank" rel="noreferrer">
              <Icon name="book" /> {urlTitle || 'Read more'}
            </a>
          )}
        </div>
      );
    }
  )
);
InfoBox.displayName = 'InfoBox';

const getInfoBoxStyles = stylesFactory((theme: GrafanaTheme, severity: AlertVariant) => {
  const color = theme.v2.palette[severity];

  return {
    wrapper: css`
      position: relative;
      padding: ${theme.spacing.md};
      background-color: ${theme.colors.bg2};
      border-left: 3px solid ${color.border};
      margin-bottom: ${theme.spacing.md};
      flex-grow: 1;
      color: ${theme.colors.textSemiWeak};
      box-shadow: ${theme.v2.shadows.z1};

      code {
        font-size: ${theme.typography.size.sm};
        background-color: ${theme.colors.bg1};
        color: ${theme.colors.text};
        border: 1px solid ${theme.colors.border2};
        border-radius: 4px;
      }

      p:last-child {
        margin-bottom: 0;
      }

      &--max-lg {
        max-width: ${theme.breakpoints.lg};
      }
    `,
    wrapperBranded: css`
      padding: ${theme.spacing.md};
      border-radius: ${theme.border.radius.md};
      position: relative;
      z-index: 0;

      &:before {
        content: '';
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        background-image: url(${theme.isLight ? panelArtLight : panelArtDark});
        border-radius: ${theme.border.radius.md};
        background-position: 50% 50%;
        background-size: cover;
        filter: saturate(80%);
        z-index: -1;
      }

      p:last-child {
        margin-bottom: 0;
      }
    `,
    docsLink: css`
      display: inline-block;
      margin-top: ${theme.spacing.md};
      font-size: ${theme.typography.size.sm};
      color: ${theme.colors.textSemiWeak};
    `,
  };
});
