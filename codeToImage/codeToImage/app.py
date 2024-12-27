from PIL import Image
from PIL import ImageDraw
from PIL import ImageFont
import argparse
import qrcode
import os

my_parser = argparse.ArgumentParser(description="Provide Token parameters")
my_parser.add_argument("--isUpdate", default=False, action="store_true",
                    help="Flag for Update profile or Create")
my_parser.add_argument(
    "-g",
    "--genratetoken",
    nargs= 2,
    metavar=("cwToken", "valueof_url"),
    help="Token cwToken & valueof_url on image QR CODE GENRATE.",)


args = my_parser.parse_args()
cwToken, valueof_url = args.genratetoken
qrcodelink= valueof_url + cwToken
fonts_path = os.path.join(os.path.dirname(os.path.dirname(__file__)), 'fonts')
# print('FONT PATH-------->',fonts_path)
font = ImageFont.truetype(os.path.join('./codeToImage/codeToImage/fonts/NotoSans-Regular.ttf'), 30)

#font = ImageFont.load_default()
if args.isUpdate:
     # print("for Update process Webp")
     img = Image.open("./codeToImage/codeToImage/EditProfile.webp")
else:
     # print("for Create process Webp")
     img = Image.open("./codeToImage/codeToImage/CreateProfile.webp")

I1 = ImageDraw.Draw(img)
I1.text((350,355), qrcodelink, fill=(4, 186, 250), font=font)
data = qrcodelink


qr = qrcode.QRCode(version=1, box_size=10, border=1)
qr.add_data(data)
qr.make(fit=True)
img_qr = qr.make_image()

#pos = (img.size[0] - img_qr.size[0], img.size[1] - img_qr.size[1])
pos = (1315,610)

img.paste(img_qr, pos)

img.save(cwToken + ".webp", to_save=True)
print(cwToken + ".webp")
