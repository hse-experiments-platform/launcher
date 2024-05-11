import pickle
import numpy as np
from sklearn.linear_model import LinearRegression
from sklearn.preprocessing import OrdinalEncoder, StandardScaler
from sklearn.model_selection import train_test_split
from sklearn.pipeline import Pipeline
from sklearn.compose import ColumnTransformer
from sklearn.impute import SimpleImputer
from sklearn.metrics import mean_squared_error
import requests
import sys
import typing
import fire
import pandas as pd

allowed_types = ['string', 'int', 'float', 'dropped', 'categorial']


# 8 5 "[('intt','int'),('floatt','float'),('categoriall','categorial'),('stringg','string')]" v2.local.Oma4GxmWWoMpYXWbO2oFEQVmNTVV8ObwS1RSFwluIl-GmvjLW7RALC1oXOHB_HEJcRbdow1BdqYGQ06XXHDZBJzC-5Wnmg7y2upHB2bR_TDjwuN2n1OKKzVvcFfi8LfrPfe0IP8hYFAvgwjDsQA3smhtB2wWKD_ajBt6zH-3NpNemFtDnUmRQISDl6VfTrItmARQKeErMEahHS01JHP1kqInbejaxxAhhkR8IxEdCIJOIU1CE3KmZZV-HV8.bnVsbA ./converted.csv localhost:8084 intt

def train(dataset_id: int, rows_count: int, schema: typing.List[typing.Tuple[str, str]], token: str,
          result_file_path: str, base_url: str, target_column: str):
    dataset_id = int(dataset_id)
    rows_count = int(rows_count)
    token = str(token)
    base_url = str(base_url)
    target_column = str(target_column)
    result_file_path = str(result_file_path)

    table = [[]]
    for t in schema:
        (name, type) = t
        if type not in allowed_types:
            print(f"type not in {allowed_types}, got {type}")
            sys.exit(3)
        table[0].append(name)

    if target_column not in table[0]:
        print(f"target column {target_column} is not present in schema")
        sys.exit(2)

    resp = requests.get(f"http://{base_url}/api/v1/datasets/{dataset_id}/rows?limit={rows_count}",
                        headers={
                            "Authorization": token})

    rows = resp.json()['rows']

    for row in rows:
        cols = row['columns']
        if len(cols) != len(table[0]):
            print(f'got invalid row, expected columns {len(table[0])}, got {len(cols)}')
        table = np.append(table, [cols], axis=0)

    df = pd.DataFrame(data=table[1:, 0:], columns=table[0, 0:])

    for (name, type) in schema:
        try:
            if type == 'int':
                df[name] = df[name].astype(int)
            elif type == 'float':
                df[name] = df[name].astype(float)
            elif type == 'categorial':
                pass
            elif type == 'string' or type == 'dropped':
                df = df.drop(name, axis=1)
        except ValueError as err:
            print(f"cannot convert column {name} to type {type}")
            sys.exit(2)

    X = df.drop([target_column], axis=1)
    Y = df[target_column]
    X_train, X_test, Y_train, Y_test = train_test_split(X, Y, test_size=0.2, random_state=44)

    s_imputer = SimpleImputer(strategy='constant', fill_value='unknown')
    o_encoder = OrdinalEncoder(handle_unknown='use_encoded_value', unknown_value=-1)
    pipe_cat = Pipeline([('imputer', s_imputer), ('encoder', o_encoder)])

    simple_imputer = SimpleImputer(strategy='median')
    std_scaler = StandardScaler()
    pipe_num = Pipeline([('imputer', simple_imputer), ('scaler', std_scaler)])

    col_transformer = ColumnTransformer(
        [('num_preproc', pipe_num, [x for x in X.columns if X[x].dtype != 'object']),
         ('cat_preproc', pipe_cat, [x for x in X.columns if X[x].dtype == 'object'])])

    reg = LinearRegression()
    final_pipe = Pipeline([('preproc', col_transformer),
                           ('model', reg)])
    # print(reg.coef_)
    final_pipe.fit(X_train, Y_train)
    preds = final_pipe.predict(X_test)
    print(mean_squared_error(Y_test, preds))
    with open(result_file_path, "wb") as file:
        pickle.dump(final_pipe, file)


if __name__ == "__main__":
    fire.Fire(train)
