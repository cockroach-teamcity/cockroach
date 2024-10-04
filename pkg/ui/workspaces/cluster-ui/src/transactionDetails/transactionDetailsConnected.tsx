// Copyright 2021 The Cockroach Authors.
//
// Use of this software is governed by the CockroachDB Software License
// included in the /LICENSE file.

import { createSelector } from "@reduxjs/toolkit";
import { connect } from "react-redux";
import { RouteComponentProps, withRouter } from "react-router-dom";
import { Dispatch } from "redux";
import { actions as localStorageActions } from "src/store/localStorage";
import { AppState, uiConfigActions } from "src/store";
import { actions as nodesActions } from "../store/nodes";
import { actions as sqlStatsActions } from "src/store/sqlStats";
import { actions as txnStatsActions } from "src/store/transactionStats";
import { TxnInsightsRequest } from "../api";
import {
  actions as transactionInsights,
  selectTxnInsightsByFingerprint,
} from "src/store/insights/transactionInsights";
import {
  TransactionDetails,
  TransactionDetailsDispatchProps,
  TransactionDetailsProps,
  TransactionDetailsStateProps,
} from "./transactionDetails";
import {
  selectTransactionsData,
  selectTransactionsLastError,
} from "../transactionsPage/transactionsPage.selectors";
import {
  selectIsTenant,
  selectHasViewActivityRedactedRole,
  selectHasAdminRole,
} from "../store/uiConfig";
import { nodeRegionsByIDSelector } from "../store/nodes";
import {
  selectTimeScale,
  selectTxnsPageLimit,
  selectTxnsPageReqSort,
} from "../store/utils/selectors";
import { StatementsRequest } from "src/api/statementsApi";
import {
  txnFingerprintIdAttr,
  getMatchParamByName,
  queryByName,
  appNamesAttr,
} from "../util";
import { TimeScale } from "../timeScaleDropdown";
import { actions as analyticsActions } from "../store/analytics";
import { selectRequestTime } from "src/transactionsPage/transactionsPage.selectors";
import { getTxnFromSqlStatsTxns } from "../transactionsPage/utils";

export const selectTransaction = createSelector(
  (state: AppState) => state.adminUI?.transactions,
  (_state: AppState, props: RouteComponentProps) => props,
  (transactionState, props) => {
    const transactions = transactionState.data?.transactions;
    if (!transactions) {
      return {
        isLoading: transactionState.inFlight,
        transaction: null,
        isValid: transactionState.valid,
      };
    }

    // We convert the empty string to null here so that ?appNames= is treated as
    // selecting all apps.
    const apps = (queryByName(props.location, appNamesAttr) || null)
      ?.split(",")
      .map(s => s.trim());

    const txnFingerprintId = getMatchParamByName(
      props.match,
      txnFingerprintIdAttr,
    );

    const transaction = getTxnFromSqlStatsTxns(
      transactions,
      txnFingerprintId,
      apps,
    );

    return {
      isLoading: transactionState.inFlight,
      transaction: transaction,
      lastUpdated: transactionState.lastUpdated,
      isValid: transactionState.valid,
    };
  },
);

const mapStateToProps = (
  state: AppState,
  props: TransactionDetailsProps,
): TransactionDetailsStateProps => {
  const { isLoading, transaction, lastUpdated, isValid } = selectTransaction(
    state,
    props,
  );
  return {
    timeScale: selectTimeScale(state),
    error: selectTransactionsLastError(state),
    isTenant: selectIsTenant(state),
    nodeRegions: nodeRegionsByIDSelector(state),
    statements: selectTransactionsData(state)?.statements,
    transaction,
    transactionFingerprintId: getMatchParamByName(
      props.match,
      txnFingerprintIdAttr,
    ),
    isLoading: isLoading,
    lastUpdated: lastUpdated,
    hasViewActivityRedactedRole: selectHasViewActivityRedactedRole(state),
    transactionInsights: selectTxnInsightsByFingerprint(state, props),
    hasAdminRole: selectHasAdminRole(state),
    isDataValid: isValid,
    limit: selectTxnsPageLimit(state),
    reqSortSetting: selectTxnsPageReqSort(state),
    requestTime: selectRequestTime(state),
  };
};

const mapDispatchToProps = (
  dispatch: Dispatch,
): TransactionDetailsDispatchProps => ({
  refreshData: (req?: StatementsRequest) =>
    dispatch(txnStatsActions.refresh(req)),
  refreshNodes: () => dispatch(nodesActions.refresh()),
  refreshUserSQLRoles: () => dispatch(uiConfigActions.refreshUserSQLRoles()),
  onTimeScaleChange: (ts: TimeScale) => {
    dispatch(
      sqlStatsActions.updateTimeScale({
        ts: ts,
      }),
    );
    dispatch(
      analyticsActions.track({
        name: "TimeScale changed",
        page: "Transaction Details",
        value: ts.key,
      }),
    );
  },
  refreshTransactionInsights: (req: TxnInsightsRequest) => {
    dispatch(transactionInsights.refresh(req));
  },
  onRequestTimeChange: (t: moment.Moment) => {
    dispatch(
      localStorageActions.update({
        key: "requestTime/StatementsPage",
        value: t,
      }),
    );
  },
});

export const TransactionDetailsPageConnected = withRouter<any, any>(
  connect(mapStateToProps, mapDispatchToProps)(TransactionDetails),
);
