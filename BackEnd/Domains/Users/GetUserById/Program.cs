using GetUserById.BusinessLogic;
using SoapCore;
namespace GetUserById
{
    public class Program
    {
        public static void Main(string[] args)
        {
            var builder = WebApplication.CreateBuilder(args);

            builder.Services.AddSoapCore();
            builder.Services.AddScoped<ISoapService, SoapService>();

            // Add services to the container.
            builder.Services.AddControllersWithViews();

            var app = builder.Build();

            app.UseRouting();
            app.UseEndpoints(endpoints =>
            {
                endpoints.UseSoapEndpoint<ISoapService>("/users.asmx", new SoapEncoderOptions(), SoapSerializer.XmlSerializer);
            });

            // Configure the HTTP request pipeline.
            if (!app.Environment.IsDevelopment())
            {
                app.UseExceptionHandler("/Home/Error");
            }
            app.UseStaticFiles();

            app.UseRouting();

            app.UseAuthorization();

            app.MapControllerRoute(
                name: "default",
                pattern: "{controller=Home}/{action=Index}/{id?}");

            app.Run();
        }
    }
}