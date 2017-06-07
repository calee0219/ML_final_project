import pandas as pd
import matplotlib.pyplot as plt
import seaborn as sns
import subprocess
import shlex

loans = pd.read_csv("./loans_data.csv")
# loans.plot(kind="scatter", x="dti", y="safe_loans")
plt.figure(figsize=(10, 16))
safe = sns.pairplot(loans[loans.safe_loans == 1], hue="safe_loans", size=10, dropna=True)
safe.savefig("safe_output.png")

unsafe = sns.pairplot(loans[loans.safe_loans == -1], hue="safe_loans", size=10, dropna=True)
unsafe.savefig("unsafe_output.png")
# sns.plt.show()

# fname='./test.pdf'
# plt.savefig(fname)
# proc=subprocess.Popen(shlex.split('lpr {f}'.format(f=fname)))