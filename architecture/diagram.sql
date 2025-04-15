                       +--------------------------+
                       |    Campaign Engine       |
                       |  (Runs per campaign)     |
                       +------------+-------------+
                                    |
                   +------------------------------+
                   |  Campaign Goroutine           |
                   |  - Checks budget              |
                   |  - Gets platform data         |
                   |  - Computes CVR/CPC           |
                   |  - Places bid                 |
                   |  - Updates spend              |
                   +------------------------------+
                                    |
                                    v
             +----------------------------------------------+
             |        Platform Simulator (shared slice)     |
             |  - Google, Meta, TikTok                      |
             |  - Updates CPC & CVR every 3 sec             |
             +----------------------------------------------+
                                    |
                                    v
                  +-----------------------------------+
                  |      CLI Console Logs            |
                  | - Platform stats (CPC, CVR)      |
                  | - Campaign bid actions           |
                  | - Budget usage                   |
                  +-----------------------------------+

