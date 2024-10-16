package awspolicy

// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

import (
	"encoding/json"
	"testing"
)

func TestPolicyEquivalence(t *testing.T) {
	cases := []struct {
		name       string
		policy1    string
		policy2    string
		equivalent bool
		err        bool
	}{
		{
			name:       "Invalid policy JSON",
			policy1:    policyTest0,
			policy2:    policyTest0,
			equivalent: false,
			err:        true,
		},

		{
			name:       "Identical policy text",
			policy1:    policyTest1,
			policy2:    policyTest1,
			equivalent: true,
		},

		{
			name:       "Action block as single item array versus string",
			policy1:    policyTest2a,
			policy2:    policyTest2b,
			equivalent: true,
		},

		{
			name:       "Action block as single item array versus string, different action",
			policy1:    policyTest3a,
			policy2:    policyTest3b,
			equivalent: false,
		},

		{
			name:       "NotAction block and ActionBlock, mixed string versus array",
			policy1:    policyTest4a,
			policy2:    policyTest4b,
			equivalent: true,
		},

		{
			name:       "NotAction block on one side",
			policy1:    policyTest5a,
			policy2:    policyTest5b,
			equivalent: false,
		},

		{
			name:       "Principal in single item array versus string",
			policy1:    policyTest6a,
			policy2:    policyTest6b,
			equivalent: true,
		},

		{
			name:       "Different principal in single item array versus string",
			policy1:    policyTest7a,
			policy2:    policyTest7b,
			equivalent: false,
		},

		{
			name:       "String principal",
			policy1:    policyTest8a,
			policy2:    policyTest8b,
			equivalent: true,
		},

		{
			name:       "String NotPrincipal",
			policy1:    policyTest9a,
			policy2:    policyTest9b,
			equivalent: true,
		},

		{
			name:       "Different NotPrincipal in single item array versus string",
			policy1:    policyTest10a,
			policy2:    policyTest10b,
			equivalent: false,
		},

		{
			name:       "Different Effect",
			policy1:    policyTest11a,
			policy2:    policyTest11b,
			equivalent: false,
		},

		{
			name:       "Different Version",
			policy1:    policyTest12a,
			policy2:    policyTest12b,
			equivalent: false,
		},

		{
			name:       "Same Condition",
			policy1:    policyTest13a,
			policy2:    policyTest13b,
			equivalent: true,
		},

		{
			name:       "Different Condition",
			policy1:    policyTest14a,
			policy2:    policyTest14b,
			equivalent: false,
		},

		{
			name:       "Condition in single string instead of array",
			policy1:    policyTest15a,
			policy2:    policyTest15b,
			equivalent: true,
		},

		{
			name:       "Multiple Condition Blocks in one policy",
			policy1:    policyTest16a,
			policy2:    policyTest16b,
			equivalent: false,
		},

		{
			name:       "Multiple Condition Blocks, same in both policies",
			policy1:    policyTest17a,
			policy2:    policyTest17b,
			equivalent: true,
		},

		{
			name:       "Multiple Statements, Equivalent",
			policy1:    policyTest18a,
			policy2:    policyTest18b,
			equivalent: true,
		},

		{
			name:       "Multiple Statements, missing one from policy 2",
			policy1:    policyTest19a,
			policy2:    policyTest19b,
			equivalent: false,
		},

		{
			name:       "Casing of Effect",
			policy1:    policyTest20a,
			policy2:    policyTest20b,
			equivalent: true,
		},

		{
			name:       "Single Statement vs []Statement",
			policy1:    policyTest21a,
			policy2:    policyTest21b,
			equivalent: true,
		},

		{
			name:       "Empty Principal set",
			policy1:    policyTest22a,
			policy2:    policyTest22b,
			equivalent: true,
		},

		{
			name:       "Empty Principals sets of different types have the same effect",
			policy1:    policyTest23a,
			policy2:    policyTest23b,
			equivalent: true,
		},

		{
			name:       "Empty Principal and missing Principal have the same effect",
			policy1:    policyTest24a,
			policy2:    policyTest24b,
			equivalent: true,
		},

		{
			name:       "Principal with empty sets and missing Principal have the same effect",
			policy1:    policyTest25a,
			policy2:    policyTest25b,
			equivalent: true,
		},

		{
			name:       "Principal with string root IAM user matches account ID",
			policy1:    policyTest26a,
			policy2:    policyTest26b,
			equivalent: true,
		},
		{
			name:       "Principal with map string root IAM user matches account ID",
			policy1:    policyTest27a,
			policy2:    policyTest27b,
			equivalent: true,
		},
		{
			name:       "Principal with map array single root IAM user matches account ID",
			policy1:    policyTest28a,
			policy2:    policyTest28b,
			equivalent: true,
		},
		{
			name:       "Principal with map array multiple root IAM user matches account ID",
			policy1:    policyTest29a,
			policy2:    policyTest29b,
			equivalent: true,
		},
		{
			name:       "Missing Statement",
			policy1:    policyTest30,
			policy2:    policyTest30,
			equivalent: true,
		},
		{
			name:       "Incorrect Statement type",
			policy1:    policyTest31,
			policy2:    policyTest31,
			equivalent: true,
			err:        false,
		},
		{
			name:       "Incorrect single Resource type",
			policy1:    policyTest32,
			policy2:    policyTest32,
			equivalent: true,
		},
		{
			name:       "Incorrect multiple Resource type",
			policy1:    policyTest33,
			policy2:    policyTest33,
			equivalent: true,
		},
		{
			name:       "Principal order not important",
			policy1:    policyTest34a,
			policy2:    policyTest34b,
			equivalent: true,
		},
		{
			name:       "Differences matter",
			policy1:    policyTest34b,
			policy2:    policyTest34c,
			equivalent: false,
		},
		{
			name:       "Boolean condition with and without quotes on single value",
			policy1:    policyTest35a,
			policy2:    policyTest35b,
			equivalent: true,
		},
		{
			name:       "Numeric condition with and without quotes on single value",
			policy1:    policyTest36a,
			policy2:    policyTest36b,
			equivalent: true,
		},
		{
			name:       "Numeric condition with and without quotes on array of values",
			policy1:    policyTest37a,
			policy2:    policyTest37b,
			equivalent: true,
		},
		{
			name:       "Condition containing empty array",
			policy1:    policyTest38a,
			policy2:    policyTest38b,
			equivalent: true,
		},
		{
			name:       "Different empty lists",
			policy1:    policyTest39a,
			policy2:    policyTest39b,
			equivalent: true,
		},
		{
			name:       "One-length lists",
			policy1:    policyTest40a,
			policy2:    policyTest40b,
			equivalent: true,
		},
		{
			name:       "Equivalent assume-role policies 1",
			policy1:    policyTest41a,
			policy2:    policyTest41b,
			equivalent: true,
		},
		{
			name:       "Equivalent assume-role policies 2",
			policy1:    policyTest42a,
			policy2:    policyTest42b,
			equivalent: true,
		},
		{
			name:       "Not equivalent assume-role policies",
			policy1:    policyTest43a,
			policy2:    policyTest43b,
			equivalent: false,
		},
		{
			name:       "Equivalence of emptiness",
			policy1:    policyTest44a,
			policy2:    policyTest44b,
			equivalent: true,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			equal, err := PoliciesAreEquivalent(tc.policy1, tc.policy2)
			if !tc.err && err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
			if tc.err && err == nil {
				t.Fatal("Expected error, none produced")
			}

			if equal != tc.equivalent {
				t.Fatalf("Bad: %s\n  Expected: %t\n       Got: %t\n", tc.name, tc.equivalent, equal)
			}
		})
	}
}

const policyTest0 = `{
  "Version": "2012-10-17",
  "Statement": [
    {
  ]
}`

const policyTest1 = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow", "Principal": {
        "Service": "spotfleet.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}`

const policyTest2a = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Principal": {
        "Service": "spotfleet.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}`

const policyTest2b = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Action": ["sts:AssumeRole"],
      "Effect": "Allow",
      "Principal": {
        "Service": "spotfleet.amazonaws.com"
      }
    }
  ]
}`

const policyTest3a = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Principal": {
        "Service": "spotfleet.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}`

const policyTest3b = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Action": ["sts:GetSessionToken"],
      "Effect": "Allow",
      "Principal": {
        "Service": "spotfleet.amazonaws.com"
      }
    }
  ]
}`

const policyTest4a = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Principal": {
        "Service": "spotfleet.amazonaws.com"
      },
      "Action": "sts:AssumeRole",
      "NotAction": ["sts:GetSessionToken"]
    }
  ]
}`

const policyTest4b = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Action": "sts:AssumeRole",
      "NotAction": "sts:GetSessionToken",
      "Effect": "Allow",
      "Principal": {
        "Service": "spotfleet.amazonaws.com"
      }
    }
  ]
}`

const policyTest5a = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Principal": {
        "Service": "spotfleet.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}`

const policyTest5b = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Action": "sts:AssumeRole",
      "NotAction": "sts:GetSessionToken",
      "Effect": "Allow",
      "Principal": {
        "Service": "spotfleet.amazonaws.com"
      }
    }
  ]
}`

const policyTest6a = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Principal": {
        "Service": "spotfleet.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}`

const policyTest6b = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Action": "sts:AssumeRole",
      "Effect": "Allow",
      "Principal": {
        "Service": ["spotfleet.amazonaws.com"]
      }
    }
  ]
}`

const policyTest7a = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Principal": {
        "Service": "ec2.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}`

const policyTest7b = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Action": "sts:AssumeRole",
      "Effect": "Allow",
      "Principal": {
        "Service": ["spotfleet.amazonaws.com"]
      }
    }
  ]
}`

const policyTest8a = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Principal": "*",
      "Action": "sts:AssumeRole"
    }
  ]
}`

const policyTest8b = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Action": "sts:AssumeRole",
      "Effect": "Allow",
      "Principal": "*"
    }
  ]
}`

const policyTest9a = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "NotPrincipal": "*",
      "Action": "sts:AssumeRole"
    }
  ]
}`

const policyTest9b = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Action": "sts:AssumeRole",
      "Effect": "Allow",
      "NotPrincipal": "*"
    }
  ]
}`

const policyTest10a = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "NotPrincipal": {
        "Service": "ec2.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}`

const policyTest10b = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Action": "sts:AssumeRole",
      "Effect": "Allow",
      "NotPrincipal": {
        "Service": ["spotfleet.amazonaws.com"]
      }
    }
  ]
}`

const policyTest11a = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Principal": {
        "Service": "spotfleet.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}`

const policyTest11b = `{
  "Version": "2012-06-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Principal": {
        "Service": "spotfleet.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}`

const policyTest12a = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Deny",
      "Principal": {
        "Service": "spotfleet.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}`

const policyTest12b = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Principal": {
        "Service": "spotfleet.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}`

const policyTest13a = `{
 "Version": "2012-10-17",
 "Statement": [
   {
     "Sid": "statement1",
     "Effect": "Allow",
     "Action": [
       "s3:PutObject"
     ],
     "Resource": [
       "arn:aws:s3:::examplebucket/*"
     ],
     "Condition": {
       "StringEquals": {
         "s3:x-amz-acl": [
           "public-read"
         ]
       }
     }
   }
 ]
}`

const policyTest13b = `{
 "Version": "2012-10-17",
 "Statement": [
   {
     "Sid": "statement1",
     "Effect": "Allow",
     "Action": [
       "s3:PutObject"
     ],
     "Resource": [
       "arn:aws:s3:::examplebucket/*"
     ],
     "Condition": {
       "StringEquals": {
         "s3:x-amz-acl": [
           "public-read"
         ]
       }
     }
   }
 ]
}`

const policyTest14a = `{
 "Version": "2012-10-17",
 "Statement": [
   {
     "Sid": "statement1",
     "Effect": "Allow",
     "Action": [
       "s3:PutObject"
     ],
     "Resource": [
       "arn:aws:s3:::examplebucket/*"
     ],
     "Condition": {
       "StringNotEquals": {
         "s3:x-amz-acl": [
           "public-read"
         ]
       }
     }
   }
 ]
}`

const policyTest14b = `{
 "Version": "2012-10-17",
 "Statement": [
   {
     "Sid": "statement1",
     "Effect": "Allow",
     "Action": [
       "s3:PutObject"
     ],
     "Resource": [
       "arn:aws:s3:::examplebucket/*"
     ],
     "Condition": {
       "StringEquals": {
         "s3:x-amz-acl": [
           "public-read"
         ]
       }
     }
   }
 ]
}`

const policyTest15a = `{
 "Version": "2012-10-17",
 "Statement": [
   {
     "Sid": "statement1",
     "Effect": "Allow",
     "Action": [
       "s3:PutObject"
     ],
     "Resource": [
       "arn:aws:s3:::examplebucket/*"
     ],
     "Condition": {
       "StringEquals": {
         "s3:x-amz-acl": "public-read"
       }
     }
   }
 ]
}`

const policyTest15b = `{
 "Version": "2012-10-17",
 "Statement": [
   {
     "Sid": "statement1",
     "Effect": "Allow",
     "Action": [
       "s3:PutObject"
     ],
     "Resource": [
       "arn:aws:s3:::examplebucket/*"
     ],
     "Condition": {
       "StringEquals": {
         "s3:x-amz-acl": [
           "public-read"
         ]
       }
     }
   }
 ]
}`

const policyTest16a = `{
 "Version": "2012-10-17",
 "Statement": [
   {
     "Sid": "statement1",
     "Effect": "Allow",
     "Action": [
       "s3:PutObject"
     ],
     "Resource": [
       "arn:aws:s3:::examplebucket/*"
     ],
     "Condition": {
       "StringEquals": {
         "s3:x-amz-acl": "public-read"
       }
     }
   }
 ]
}`

const policyTest16b = `{
 "Version": "2012-10-17",
 "Statement": [
   {
     "Sid": "statement1",
     "Effect": "Allow",
     "Action": [
       "s3:PutObject"
     ],
     "Resource": [
       "arn:aws:s3:::examplebucket/*"
     ],
     "Condition" :  {
       "DateGreaterThan" : {
         "aws:CurrentTime" : "2013-08-16T12:00:00Z"
       },
       "DateLessThan": {
         "aws:CurrentTime" : "2013-08-16T15:00:00Z"
       },
       "IpAddress" : {
         "aws:SourceIp" : ["192.0.2.0/24", "203.0.113.0/24"]
       }
     }
   }
 ]
}`

const policyTest17a = `{
 "Version": "2012-10-17",
 "Statement": [
   {
     "Sid": "statement1",
     "Effect": "Allow",
     "Action": [
       "s3:PutObject"
     ],
     "Resource": [
       "arn:aws:s3:::examplebucket/*"
     ],
     "Condition" :  {
       "DateGreaterThan" : {
         "aws:CurrentTime" : "2013-08-16T12:00:00Z"
       },
       "DateLessThan": {
         "aws:CurrentTime" : "2013-08-16T15:00:00Z"
       },
       "IpAddress" : {
         "aws:SourceIp" : ["192.0.2.0/24", "203.0.113.0/24"]
       }
     }
   }
 ]
}`

const policyTest17b = `{
 "Version": "2012-10-17",
 "Statement": [
   {
     "Sid": "statement1",
     "Effect": "Allow",
     "Action": [
       "s3:PutObject"
     ],
     "Resource": [
       "arn:aws:s3:::examplebucket/*"
     ],
     "Condition" :  {
       "DateGreaterThan" : {
         "aws:CurrentTime" : "2013-08-16T12:00:00Z"
       },
       "DateLessThan": {
         "aws:CurrentTime" : "2013-08-16T15:00:00Z"
       },
       "IpAddress" : {
         "aws:SourceIp" : ["192.0.2.0/24", "203.0.113.0/24"]
       }
     }
   }
 ]
}`

const policyTest18a = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "s3:ListAllMyBuckets",
        "s3:GetBucketLocation"
      ],
      "Resource": "arn:aws:s3:::*"
    },
    {
      "Effect": "Allow",
      "Action": "s3:ListBucket",
      "Resource": "arn:aws:s3:::BUCKET-NAME",
      "Condition": {"StringLike": {"s3:prefix": [
        "",
        "home/",
        "home/${aws:username}/"
      ]}}
    },
    {
      "Effect": "Allow",
      "Action": "s3:*",
      "Resource": [
        "arn:aws:s3:::BUCKET-NAME/home/${aws:username}",
        "arn:aws:s3:::BUCKET-NAME/home/${aws:username}/*"
      ]
    }
  ]
}`

const policyTest18b = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "s3:ListAllMyBuckets",
        "s3:GetBucketLocation"
      ],
      "Resource": "arn:aws:s3:::*"
    },
    {
      "Effect": "Allow",
      "Action": "s3:ListBucket",
      "Resource": "arn:aws:s3:::BUCKET-NAME",
      "Condition": {"StringLike": {"s3:prefix": [
        "",
        "home/",
        "home/${aws:username}/"
      ]}}
    },
    {
      "Effect": "Allow",
      "Action": "s3:*",
      "Resource": [
        "arn:aws:s3:::BUCKET-NAME/home/${aws:username}",
        "arn:aws:s3:::BUCKET-NAME/home/${aws:username}/*"
      ]
    }
  ]
}`

const policyTest19a = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "s3:ListAllMyBuckets",
        "s3:GetBucketLocation"
      ],
      "Resource": "arn:aws:s3:::*"
    },
    {
      "Effect": "Allow",
      "Action": "s3:*",
      "Resource": [
        "arn:aws:s3:::BUCKET-NAME/home/${aws:username}",
        "arn:aws:s3:::BUCKET-NAME/home/${aws:username}/*"
      ]
    },
    {
      "Effect": "Allow",
      "Action": "s3:ListBucket",
      "Resource": "arn:aws:s3:::BUCKET-NAME",
      "Condition": {"StringLike": {"s3:prefix": [
        "",
        "home/",
        "home/${aws:username}/"
      ]}}
    }
  ]
}`

const policyTest19b = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "s3:ListAllMyBuckets",
        "s3:GetBucketLocation"
      ],
      "Resource": "arn:aws:s3:::*"
    },
    {
      "Effect": "Allow",
      "Action": "s3:*",
      "Resource": [
        "arn:aws:s3:::BUCKET-NAME/home/${aws:username}",
        "arn:aws:s3:::BUCKET-NAME/home/${aws:username}/*"
      ]
    }
  ]
}`

const policyTest20a = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Principal": {
        "Service": "spotfleet.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
  ]
}`

const policyTest20b = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Action": ["sts:AssumeRole"],
      "Effect": "allow",
      "Principal": {
        "Service": "spotfleet.amazonaws.com"
      }
    }
  ]
}`

const policyTest21a = `{
  "Version": "2012-10-17",
  "Statement":
    {
      "Sid": "",
      "Effect": "Allow",
      "Principal": {
        "Service": "spotfleet.amazonaws.com"
      },
      "Action": "sts:AssumeRole"
    }
}`

const policyTest21b = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Action": ["sts:AssumeRole"],
      "Effect": "allow",
      "Principal": {
        "Service": "spotfleet.amazonaws.com"
      }
    }
  ]
}`

const policyTest22a = `{
  "Version": "2012-10-17",
  "Statement":
    {
      "Sid": "",
      "Effect": "Allow",
      "Principal": {},
      "Action": "sts:AssumeRole"
    }
}`

const policyTest22b = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Action": ["sts:AssumeRole"],
      "Effect": "allow",
      "Principal": {
        "Service": []
      }
    }
  ]
}`

const policyTest23a = `{
  "Version": "2012-10-17",
  "Statement":
    {
      "Sid": "",
      "Effect": "Allow",
      "Principal": {
        "Service": []
      },
      "Action": "sts:AssumeRole"
    }
}`

const policyTest23b = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Action": ["sts:AssumeRole"],
      "Effect": "allow",
      "Principal": {
        "AWS": []
      }
    }
  ]
}`

const policyTest24a = `{
  "Version": "2012-10-17",
  "Statement":
    {
      "Sid": "",
      "Effect": "Allow",
      "Principal": {},
      "Action": "sts:AssumeRole"
    }
}`

const policyTest24b = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Action": ["sts:AssumeRole"],
      "Effect": "allow"
    }
  ]
}`

const policyTest25a = `{
  "Version": "2012-10-17",
  "Statement":
    {
      "Sid": "",
      "Effect": "Allow",
      "Principal": {
        "Service": [],
        "AWS": []
      },
      "Action": "sts:AssumeRole"
    }
}`

const policyTest25b = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Action": ["sts:AssumeRole"],
      "Effect": "allow"
    }
  ]
}`

const policyTest26a = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Action": "*",
      "Resource": "*",
      "Principal": "123456789012"
    }
  ]
}`

const policyTest26b = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Action": "*",
      "Resource": "*",
      "Principal": "arn:PARTITION:iam::123456789012:root"
    }
  ]
}`

const policyTest27a = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Action": "*",
      "Resource": "*",
      "Principal": {
        "AWS": "123456789012"
      }
    }
  ]
}`

const policyTest27b = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Action": "*",
      "Resource": "*",
      "Principal": {
        "AWS": "arn:PARTITION:iam::123456789012:root"
      }
    }
  ]
}`

const policyTest28a = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Action": "*",
      "Resource": "*",
      "Principal": {
        "AWS": [
          "123456789012"
        ]
      }
    }
  ]
}`

const policyTest28b = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Action": "*",
      "Resource": "*",
      "Principal": {
        "AWS": [
          "arn:PARTITION:iam::123456789012:root"
        ]
      }
    }
  ]
}`

const policyTest29a = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Action": "*",
      "Resource": "*",
      "Principal": {
        "AWS": [
          "123456789012",
          "arn:PARTITION:iam::999999999999:root"
        ]
      }
    }
  ]
}`

const policyTest29b = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Action": "*",
      "Resource": "*",
      "Principal": {
        "AWS": [
          "arn:PARTITION:iam::123456789012:root",
          "arn:PARTITION:iam::999999999999:root"
        ]
      }
    }
  ]
}`

const policyTest30 = `{
  "Version": "2012-10-17"
}`

const policyTest31 = `{
  "Version": "2012-10-17",
  "Statement": 42
}`

const policyTest32 = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "statement1",
      "Effect": "Allow",
      "Action": [
        "s3:PutObject"
      ],
      "Resource": {}
    }
  ]
}`

const policyTest33 = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "statement1",
      "Effect": "Allow",
      "Action": [
        "s3:PutObject"
      ],
      "Resource": [[42]]
    }
  ]
}`

const policyTest34a = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Action": "*",
      "Resource": "*",
      "Principal": {
        "AWS": [
          "arn:aws:iam::123456789012:role/franz-ferdinand",
          "arn:aws:iam::123456789012:role/morgan-page",
          "arn:aws:iam::123456789012:role/built-by-titan",
          "arn:aws:iam::123456789012:role/kristina_sky",
          "arn:aws:iam::123456789012:role/maria_becerra",
          "arn:aws:iam::123456789012:role/jodie_knight"
        ]
      }
    }
  ]
}`

const policyTest34b = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Action": "*",
      "Resource": "*",
      "Principal": {
        "AWS": [
          "arn:aws:iam::123456789012:role/maria_becerra",
          "arn:aws:iam::123456789012:role/morgan-page",
          "arn:aws:iam::123456789012:role/kristina_sky",
          "arn:aws:iam::123456789012:role/jodie_knight",
          "arn:aws:iam::123456789012:role/franz-ferdinand",
          "arn:aws:iam::123456789012:role/built-by-titan"
        ]
      }
    }
  ]
}`

const policyTest34c = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "",
      "Effect": "Allow",
      "Action": "*",
      "Resource": "*",
      "Principal": {
        "AWS": [
          "arn:aws:iam::123456789012:role/maria_becerra",
          "arn:aws:iam::123456789012:role/morgan-page",
          "arn:aws:iam::123456789012:role/kristina_sky",
          "arn:aws:iam::123456789012:role/jodie_knight",
          "arn:aws:iam::123456789012:role/franz-ferdinand",
          "arn:aws:iam::123456789012:role/idina-menzel"
        ]
      }
    }
  ]
}`

const policyTest35a = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "statement1",
      "Effect": "Allow",
      "Action": [
        "s3:PutObject"
      ],
      "Resource": [
        "arn:aws:s3:::examplebucket/*"
      ],
      "Condition": {
        "Bool": {
          "aws:MultiFactorAuthPresent": true
        }
      }
    }
  ]
 }`

const policyTest35b = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "statement1",
      "Effect": "Allow",
      "Action": [
        "s3:PutObject"
      ],
      "Resource": [
        "arn:aws:s3:::examplebucket/*"
      ],
      "Condition": {
        "Bool": {
          "aws:MultiFactorAuthPresent": "true"
        }
      }
    }
  ]
 }`

const policyTest36a = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "statement1",
      "Effect": "Allow",
      "Action": [
        "s3:PutObject"
      ],
      "Resource": [
        "arn:aws:s3:::examplebucket/*"
      ],
      "Condition": {
        "NumericLessThanEquals": {
          "aws:MultiFactorAuthAge": 100
        }
      }
    }
  ]
 }`

const policyTest36b = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "statement1",
      "Effect": "Allow",
      "Action": [
        "s3:PutObject"
      ],
      "Resource": [
        "arn:aws:s3:::examplebucket/*"
      ],
      "Condition": {
        "NumericLessThanEquals": {
          "aws:MultiFactorAuthAge": "100"
        }
      }
    }
  ]
 }`

const policyTest37a = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "statement1",
      "Effect": "Allow",
      "Action": [
        "s3:PutObject"
      ],
      "Resource": [
        "arn:aws:s3:::examplebucket/*"
      ],
      "Condition": {
        "NumericEquals": {
          "aws:MultiFactorAuthAge": [100.01, 200.2]
        }
      }
    }
  ]
 }`

const policyTest37b = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "statement1",
      "Effect": "Allow",
      "Action": [
        "s3:PutObject"
      ],
      "Resource": [
        "arn:aws:s3:::examplebucket/*"
      ],
      "Condition": {
        "NumericEquals": {
          "aws:MultiFactorAuthAge": ["100.01", "200.2"]
        }
      }
    }
  ]
 }`

const policyTest38a = `{
  "Id":"sqspolicy",
  "Statement":[
     {
        "Action":"sqs:SendMessage",
        "Condition":{
           "ArnEquals":{
              "aws:SourceArn":[]
           }
        },
        "Effect":"Allow",
        "Principal":"*",
        "Resource":"arn:aws:sqs:***:redacted:redacted",
        "Sid":"AllowSNSSubscriptions"
     }
  ],
  "Version":"2012-10-17"
}`

const policyTest38b = `{
  "Version":"2012-10-17",
  "Id":"sqspolicy",
  "Statement":[
     {
        "Sid":"AllowSNSSubscriptions",
        "Effect":"Allow",
        "Principal":"*",
        "Action":"sqs:SendMessage",
        "Resource":"arn:aws:sqs:***:redacted:redacted",
        "Condition":{
           "ArnEquals":{
              "aws:SourceArn":[]
           }
        }
     }
  ]
}`

const policyTest39a = `[]`

const policyTest39b = `[{

}]`

const policyTest40a = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "statement1",
      "Effect": "Allow",
      "Action": [
        "s3:PutObject"
      ],
      "Resource": [42, 86]
    }
  ]
}`

const policyTest40b = `{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Sid": "statement1",
      "Effect": "Allow",
      "Action": "s3:PutObject",
      "Resource": [86, 42]
    }
  ]
}`

const policyTest41a = `{"Statement":[{"Action":"sts:AssumeRole","Effect":"Allow","Principal":{"Service":"ec2.amazonaws.com"}}],"Version":"2012-10-17"}`
const policyTest41b = `{"Statement":[{"Action":["sts:AssumeRole"],"Effect":"Allow","Principal":{"Service":["ec2.amazonaws.com"]}}],"Version":"2012-10-17"}`

const policyTest42a = `{"Statement":[{"Action":"sts:AssumeRole","Effect":"Allow","Principal":{"Service":"ec2.amazonaws.com"}}],"Version":"2012-10-17"}`
const policyTest42b = `{"Statement":[{"Action":"sts:AssumeRole","Effect":"Allow","Principal":{"Service":["ec2.amazonaws.com"]}}],"Version":"2012-10-17"}`

const policyTest43a = `{"Statement":[{"Action":"sts:AssumeRole","Effect":"Allow","Principal":{"Service":"ec2.amazonaws.com"}}],"Version":"2012-10-17"}`
const policyTest43b = `{"Statement":[{"Action":"sts:AssumeRole","Effect":"Allow","Principal":{"Service":["rds.amazonaws.com"]}}],"Version":"2012-10-17"}`

const policyTest44a = ``
const policyTest44b = `{}`

func TestStringValueSlicesEqualIgnoreOrder(t *testing.T) {
	equal := []interface{}{
		[]interface{}{
			[]string{"a", "b", "c"},
			[]string{"a", "b", "c"},
		},
		[]interface{}{
			[]string{"b", "a", "c"},
			[]string{"a", "b", "c"},
		},
		[]interface{}{
			[]string{"apple", "carrot", "tomato"},
			[]string{"tomato", "apple", "carrot"},
		},
		[]interface{}{
			[]string{"Application", "Barrier", "Chilly", "Donut"},
			[]string{"Barrier", "Application", "Donut", "Chilly"},
		},
	}
	for _, v := range equal {
		if !stringSlicesEqualIgnoreOrder(v.([]interface{})[0].([]string), v.([]interface{})[1].([]string)) {
			t.Fatalf("%v should be equal: %v", v.([]interface{})[0].([]string), v.([]interface{})[1].([]string))
		}
	}

	notEqual := []interface{}{
		[]interface{}{
			[]string{"c", "b", "c"},
			[]string{"a", "b", "c"},
		},
		[]interface{}{
			[]string{"b", "a", "c"},
			[]string{"a", "bread", "c"},
		},
		[]interface{}{
			[]string{"apple", "carrot", "tomato"},
			[]string{"tomato", "apple"},
		},
		[]interface{}{
			[]string{"Application", "Barrier", "Chilly", "Donut"},
			[]string{"Barrier", "Applications", "Donut", "Chilly"},
		},
		[]interface{}{
			[]string{},
			[]string{"Barrier", "Applications", "Donut", "Chilly"},
		},
	}
	for _, v := range notEqual {
		if stringSlicesEqualIgnoreOrder(v.([]interface{})[0].([]string), v.([]interface{})[1].([]string)) {
			t.Fatalf("%v should not be equal: %v", v.([]interface{})[0].([]string), v.([]interface{})[1].([]string))
		}
	}
}

func TestIntermediatePolicyDocument(t *testing.T) {
	cases := []struct {
		name                   string
		inputPolicy            string
		expectedPolicyDocument *policyDocument
		err                    bool
	}{
		{
			name: "invalid policy with null string statement",
			inputPolicy: `{
				"Version": "2012-10-17",
				"Statement": "null"
			}`,
			err: true,
		},
		{
			name: "policy with version and nil statement",
			inputPolicy: `{
				"Version": "2012-10-17",
				"Statement": null
			}`,
			expectedPolicyDocument: &policyDocument{
				Version:    "2012-10-17",
				Statements: nil,
			},
		},
		{
			name:        "basic policy",
			inputPolicy: policyTest1,
			expectedPolicyDocument: &policyDocument{
				Version: "2012-10-17",
				Statements: []*policyStatement{
					{
						Sid:    "",
						Effect: "Allow",
						Principals: map[string]interface{}{
							"Service": "spotfleet.amazonaws.com",
						},
						Actions: "sts:AssumeRole",
					},
				},
			},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			policy1intermediate := &intermediatePolicyDocument{}
			err := json.Unmarshal([]byte(tc.inputPolicy), policy1intermediate)
			if err != nil {
				t.Fatalf("Error unmarshaling policy: %s", err)
			}

			actual, err := policy1intermediate.document()

			if !tc.err && err != nil {
				t.Fatalf("Unexpected error: %s", err)
			}
			if tc.err && err == nil {
				t.Fatal("Expected error, none produced")
			}

			if !actual.equals(tc.expectedPolicyDocument) {
				t.Fatalf("Bad: %s\n  Expected: %v\n       Got: %v\n", tc.name, tc.expectedPolicyDocument, actual)
			}
		})
	}
}
