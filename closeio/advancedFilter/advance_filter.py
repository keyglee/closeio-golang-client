"""
{
    "include_counts": True,
    "_limit": query_limit,
    "query": {
        "negate": False,
        "queries": [
            {
                "negate": False,
                "object_type": "lead",
                "type": "object_type"
            },
            {
                "negate": False,
                "queries": [
                    {
                        "negate": False,
                        "queries": [
                            {
                                "condition": {
                                    "before": {
                                        "type": "now"
                                    },
                                    "on_or_after": {
                                        "direction": "past",
                                        "moment": {
                                            "type": "now"
                                        },
                                        "offset": {
                                            "days": 0,
                                            "hours": 50,
                                            "minutes": 0,
                                            "months": 0,
                                            "seconds": 0,
                                            "weeks": 0,
                                            "years": 0
                                        },
                                        "type": "offset",
                                        "which_day_end": "start"
                                    },
                                    "type": "moment_range"
                                },
                                "field": {
                                    "field_name": "date_updated",
                                    "object_type": "lead",
                                    "type": "regular_field"
                                },
                                "negate": False,
                                "type": "field_condition"
                            },
                            {
                                "condition": {
                                    "gt": 0,
                                    "type": "number_range"
                                },
                                "field": {
                                    "field_name": "num_email_addresses",
                                    "object_type": "lead",
                                    "type": "regular_field"
                                },
                                "negate": False,
                                "type": "field_condition"
                            },
                            {
                                "condition": {
                                    "type": "exists"
                                },
                                "field": {
                                    "custom_field_id": "self.custom_subscription_status",
                                    "type": "custom_field"
                                },
                                "negate": False,
                                "type": "field_condition"
                            },
                            {
                                "condition": {
                                    "type": "exists"
                                },
                                "field": {
                                    "custom_field_id": self.custom_market,
                                    "type": "custom_field"
                                },
                                "negate": False,
                                "type": "field_condition"
                            },
                            {
                                "condition": {
                                    "type": "exists"
                                },
                                "field": {
                                    "custom_field_id": self.custom_dispo_lead_owner,
                                    "type": "custom_field"
                                },
                                "negate": False,
                                "type": "field_condition"
                            },
                            {
                                "condition": {
                                    "type": "exists"
                                },
                                "field": {
                                    "custom_field_id": self.custom_dispo_type,
                                    "type": "custom_field"
                                },
                                "negate": False,
                                "type": "field_condition"
                            }
                        ],
                        "type": "and"
                    }
                ],
                "type": "and"
            }
        ],
        "type": "and"
    },
    "_fields": {
        "lead": ["id", "name", "title", "contacts", "custom"]
    }
}
"""