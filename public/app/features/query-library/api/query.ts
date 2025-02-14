import { BaseQueryFn } from '@reduxjs/toolkit/query/react';
import { lastValueFrom } from 'rxjs';

import { BackendSrvRequest, getBackendSrv, isFetchError } from '@grafana/runtime/src/services/backendSrv';

import { getAPINamespace } from '../../../api/utils';

/**
 * @alpha
 */
export const API_VERSION = 'peakq.grafana.app/v0alpha1';

/**
 * @alpha
 */
export enum QueryTemplateKinds {
  QueryTemplate = 'QueryTemplate',
}

/**
 * Query Library is an experimental feature. API (including the URL path) will likely change.
 *
 * @alpha
 */
export const BASE_URL = `/apis/${API_VERSION}/namespaces/${getAPINamespace()}`;

interface QueryLibraryBackendRequest extends BackendSrvRequest {
  body?: BackendSrvRequest['data'];
}

export const baseQuery: BaseQueryFn<QueryLibraryBackendRequest, unknown, Error> = async (requestOptions) => {
  try {
    const responseObservable = getBackendSrv().fetch({
      url: `${BASE_URL}/${requestOptions.url ?? ''}`,
      showErrorAlert: true,
      method: requestOptions.method || 'GET',
      data: requestOptions.body,
      params: requestOptions.params,
      headers: { ...requestOptions.headers },
    });
    return await lastValueFrom(responseObservable);
  } catch (error) {
    if (isFetchError(error)) {
      return { error: new Error(error.data.message) };
    } else if (error instanceof Error) {
      return { error };
    } else {
      return { error: new Error('Unknown error') };
    }
  }
};
