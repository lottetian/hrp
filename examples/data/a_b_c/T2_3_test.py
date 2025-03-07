# NOTE: Generated By HttpRunner v4.0.0-beta
# FROM: a-b.c/2 3.yml


import sys
from pathlib import Path

sys.path.insert(0, str(Path(__file__).parent.parent))


from httprunner import HttpRunner, Config, Step, RunRequest, RunTestCase

from a_b_c.T1_test import TestCaseT1 as T1


class TestCaseT23(HttpRunner):

    config = (
        Config("reference testcase unittest for abnormal folder path")
        .base_url("https://postman-echo.com")
        .verify(False)
    )

    teststeps = [
        Step(RunTestCase("request with functions").call(T1).export(*["session_foo2"])),
        Step(
            RunRequest("post form data")
            .with_variables(**{"foo1": "bar12"})
            .post("/post")
            .with_headers(
                **{
                    "User-Agent": "HttpRunner/${get_httprunner_version()}",
                    "Content-Type": "application/x-www-form-urlencoded",
                }
            )
            .with_data("foo1=$foo1&foo2=$session_foo2")
            .validate()
            .assert_equal("status_code", 200)
            .assert_equal("body.form.foo1", "bar12")
            .assert_equal("body.form.foo2", "config_bar2")
        ),
    ]


if __name__ == "__main__":
    TestCaseT23().test_start()
